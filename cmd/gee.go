package cmd

import (
	"github.com/daishan/go7/gee"
)

func Start() {
	engine := gee.NewEngine()
	engine.Get("/", geeHandler)
	engine.Run("localhost:9000")
}

func geeHandler(c *gee.Context) {
	c.Json(200, gee.H{"message": "hello gee", "status": "success"})
}
