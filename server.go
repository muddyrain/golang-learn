package main

import "net/http"

type Server interface {
	Route(pattern string, handler http.HandlerFunc)
	Start(address string) error
}
type sdkHttpServer struct {
	Name string
}

type Header map[string][]string
