package gee

import (
	"fmt"
	"testing"
)

func NewRouters() *router {
	r := NewRouter()
	r.AddRoute("Get", "/", func(c *Context) {
		fmt.Println("root directory")
	})
	r.AddRoute("Get", "/hello/:name", func(c *Context) {
		fmt.Println("hello name = ", c.Param("name"))
	})

	r.AddRoute("get", "/hello/:age/:sex", func(c *Context) {
		fmt.Println("hello age = ", c.Param("age"), " sex = ", c.Param("sex"))
	})

	r.AddRoute("Get", "/hello/*filepath", func(c *Context) {
		fmt.Println("hello filepath = ", c.Param("filepath"))
	})
	return r
}

func TestGetRoute(t *testing.T) {
	r := NewRouters()

	// n, params := r.getRoute("Get", "/")
	// if n == nil {
	// 	t.Fatal("root match failed")
	// }
	// fmt.Println("pattern = ", n.pattern, " params = ", params)

	// n, params = r.getRoute("Get", "/hello/ds")
	// if n == nil {
	// 	t.Fatal("name / age match failed.")
	// }
	// fmt.Println("pattern = ", n.pattern, " params = ", params)

	n, params := r.getRoute("Get", "/hello/29/male")
	if n == nil {
		t.Fatal("hello age male match failed.")
	}
	fmt.Println("pattern = ", n.pattern, " params = ", params)

	n, params = r.getRoute("Get", "/hello/29/male/danggang")
	if n == nil {
		t.Fatal("hello age male danggang match failed.")
	}
	fmt.Println("pattern = ", n.pattern, " params = ", params)
}
