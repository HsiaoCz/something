package slick

import "log"

type RouterGroup struct {
	prefix     string
	middleware []Handlerfunc
	parent     *RouterGroup
	slick      *Slick
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	slick := g.slick
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		slick:  slick,
	}
	slick.groups = append(slick.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler Handlerfunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.slick.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (g *RouterGroup) GET(pattern string, handler Handlerfunc) {
	g.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (g *RouterGroup) POST(pattern string, handler Handlerfunc) {
	g.addRoute("POST", pattern, handler)
}
