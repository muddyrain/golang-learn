package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	_, err := w.Write([]byte("check your terminal"))
	if err != nil {
		return
	}
}
func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
