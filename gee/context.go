package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusType int
type H map[string]interface{}

const (
	STATUSCODE_OK       StatusType = 200
	STATUSCODE_NotFound StatusType = 404
)

type Context struct {
	Req    http.Request
	Writer http.ResponseWriter

	Method string
	Path   string

	StatusCode StatusType
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetStatus(code StatusType) {
	c.StatusCode = code
	c.Writer.WriteHeader(int(code))
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code StatusType, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(code StatusType, obj interface{}) {
	c.SetHeader("Content-Type", "text/json")
	c.SetStatus(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}

}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:    *r,
		Writer: w,
		Method: r.Method,
		Path:   r.URL.Path,
	}
}
