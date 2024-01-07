package infra

import (
	"io"
	"net/http"
)

type Retriever struct {
}

func (r Retriever) Get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// ReadAll 被废弃了，使用 io.ReadAll
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
