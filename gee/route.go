package gee

import (
	"log"
	"strings"
)

type HandlerFunc func(c *Context)

type router struct {
	handlers map[string]HandlerFunc
	routes   map[string]*tri_node
}

func (*router) parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts

}

func (r *router) AddRoute(method string, pattern string, handler HandlerFunc) {
	if handler == nil {
		log.Fatal("handler is nil")
		return
	}
	key := r.getKey(method, pattern)
	parts := r.parsePattern(pattern)

	r.handlers[key] = handler
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = &tri_node{
			part:     "/",
			children: make([]*tri_node, 0),
		}
	}
	r.routes[method].insert(pattern, parts, 0)
}

func (r *router) getRoute(method string, path string) (*tri_node, map[string]string) {
	parts := r.parsePattern(path)
	params := make(map[string]string)

	result := r.routes[method].search(parts, 0)
	if result != nil {
		patternParts := r.parsePattern(result.pattern)
		for i, part := range patternParts {
			if part[0] == ':' {
				params[part[1:]] = parts[i]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(parts[i:], "/")
			}
		}
		return result, params
	}

	return nil, nil
}

func (r *router) Handle(c *Context) {
	key := r.getKey(c.Method, c.Path)
	_, params := r.getRoute(c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		c.Params = params
		handler(c)
	} else {
		c.Writer.WriteHeader(404)
		_, _ = c.Writer.Write([]byte("404  NOT FOUND: " + c.Path))
	}
}

func (r *router) getKey(method string, pattern string) string {
	return method + "-" + pattern
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		routes:   make(map[string]*tri_node),
	}
}
