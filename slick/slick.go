package slick

import (
	"net/http"
)

type Handlerfunc func(*Context)

type Slick struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Slick {
	slick := &Slick{
		router: newRouter(),
	}
	slick.RouterGroup = &RouterGroup{slick: slick}
	slick.groups = []*RouterGroup{slick.RouterGroup}
	return slick
}

func (s *Slick) addRoute(method string, pattern string, handler Handlerfunc) {
	s.router.addRoute(method, pattern, handler)
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

func (s *Slick) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	s.router.handle(c)
}
