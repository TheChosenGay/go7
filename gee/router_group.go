package gee

// 对相同前缀的路由进行分组管理
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parentGroup *RouterGroup
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix:      group.prefix + prefix,
		engine:      group.engine,
		parentGroup: group,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.rout.AddRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
