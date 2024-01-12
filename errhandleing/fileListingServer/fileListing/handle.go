package fileListing

import (
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type UserError string

func (u UserError) Error() string {
	return u.Message()
}

func (u UserError) Message() string {
	return string(u)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return UserError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	res, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(res)
	return nil
}
