package gee

import (
	"net/http"
)

type Engine struct {
	*RouterGroup
	rout   *router
	groups []*RouterGroup
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.rout.AddRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.rout.AddRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func NewEngine() *Engine {
	engine := &Engine{
		rout: NewRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 统一拦截
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	e.rout.Handle(context)
}
