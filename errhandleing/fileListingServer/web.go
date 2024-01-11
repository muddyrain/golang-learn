package main

import (
	"awesomeProject/errhandleing/fileListingServer/fileListing"
	"fmt"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			fmt.Println("Error handling request: ", err.Error())
			code := http.StatusOK
			if os.IsNotExist(err) {
				code = http.StatusNotFound
			} else if os.IsPermission(err) {
				code = http.StatusForbidden
			} else {
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(fileListing.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
