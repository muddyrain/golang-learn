package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}
func (e testingUserError) Message() string {
	return string(e)
}
func errUserErr(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrInvalid
}
func noError(writer http.ResponseWriter, request *http.Request) error {
	writer.Write([]byte("no error"))
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, http.StatusInternalServerError, "Internal Server Error"},
	{errUserErr, http.StatusBadRequest, "user error"},
	{errNotFound, http.StatusNotFound, "Not Found"},
	{errNoPermission, http.StatusForbidden, "Forbidden"},
	{errUnknown, http.StatusInternalServerError, "Internal Server Error"},
	{noError, http.StatusOK, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, code int, message string, t *testing.T) {
	b, _ := io.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != code || message != body {
		t.Errorf("expect (%d, %s); got (%d, %s)", code, message, resp.StatusCode, body)
	}
}
