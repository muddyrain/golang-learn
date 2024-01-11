package fileListing

import (
	"io"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
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
