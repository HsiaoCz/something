package slick

import (
	"net/http"
	"strings"
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
	var middlewares []Handlerfunc
	for _, group := range s.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	s.router.handle(c)
}
