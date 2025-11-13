package cmd

import (
	"github.com/daishan/go7/gee"
)

func Start() {
	engine := gee.NewEngine()
	engine.Get("/", geeHandler)
	engine.Get("/query", queryHandler)
	engine.Run("localhost:9000")
}

func queryHandler(c *gee.Context) {
	c.String(200, "query name = %s", c.Query("name"))
}

func geeHandler(c *gee.Context) {
	c.Json(200, gee.H{"message": "hello gee", "status": "success"})
}
