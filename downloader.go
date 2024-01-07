package main

import (
	"awesomeProject/infra"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func getRetriever() Retriever {
	return infra.Retriever{}
}
func main() {
	retriever := getRetriever()
	fmt.Printf("%s\n", retriever.Get("https://www.imooc.com"))
}
