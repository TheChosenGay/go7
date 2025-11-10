package gee

import (
	"log"
)

type HandlerFunc func(c *Context)

type router struct {
	routes map[string]HandlerFunc
}

func (r *router) AddRoute(method string, pattern string, handler HandlerFunc) {
	if handler == nil {
		log.Fatal("handler is nil")
		return
	}
	key := r.getKey(method, pattern)
	r.routes[key] = handler
}

func (r *router) Handle(c *Context) {
	key := r.getKey(c.Method, c.Path)
	if handler, ok := r.routes[key]; ok {
		handler(c)
	} else {
		c.Writer.WriteHeader(404)
		_, _ = c.Writer.Write([]byte("404 NOT FOUND: " + c.Path))
	}
}

func (r *router) getKey(method string, pattern string) string {
	return method + "-" + pattern
}

func NewRouter() *router {
	return &router{
		routes: make(map[string]HandlerFunc),
	}
}
