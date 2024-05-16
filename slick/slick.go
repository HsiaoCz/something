package slick

import (
	"fmt"
	"net/http"
)

type Handlerfunc func(w http.ResponseWriter, r *http.Request)

type Slick struct {
	router map[string]Handlerfunc
}

func New() *Slick {
	return &Slick{
		router: make(map[string]Handlerfunc),
	}
}

func (s *Slick) addRoute(method string, pattern string, handler Handlerfunc) {
	key := method + "-" + pattern
	s.router[key] = handler
}

// GET defines the method to add GET request
func (s *Slick) GET(pattern string, handler Handlerfunc) {
	s.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (s *Slick) POST(pattern string, handler Handlerfunc) {
	s.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (s *Slick) Run(addr string) (err error) {
	return http.ListenAndServe(addr, s)
}

func (s *Slick) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := s.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
