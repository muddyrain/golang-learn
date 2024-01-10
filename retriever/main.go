package main

import (
	"awesomeProject/retriever/mock"
	really "awesomeProject/retriever/real"
	"fmt"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

//	func download(r Retriever) string {
//		return r.Get(url)
//	}
//func post(p Poster) {
//	res := p.Post(url, map[string]string{
//		"name":   "ccmouse",
//		"course": "golang",
//	})
//	fmt.Println(res)
//}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"Contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Println("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *really.Retriever:
		{
			fmt.Println("UserAgent:", v.UserAgent)
			fmt.Println("Timeout:", v.Timeout)
		}
	}
	fmt.Println("-------------")
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &really.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	// Type assertion
	reallyRetriever := r.(*really.Retriever)
	fmt.Println("reallyRetriever", reallyRetriever.UserAgent)
	fmt.Println("reallyRetriever", reallyRetriever.Timeout)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//fmt.Printf("%s\n", download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
