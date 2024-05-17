package slick

import (
	"log"
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	middlewares []Handlerfunc
	parent      *RouterGroup
	slick       *Slick
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

func (g *RouterGroup) Use(middleware ...Handlerfunc) {
	g.middlewares = append(g.middlewares, middleware...)
}

func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) Handlerfunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.W, c.R)
	}
}

func (g *RouterGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	g.GET(urlPattern, handler)
}
