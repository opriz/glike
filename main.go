package main

import (
	"gin-like/glike"
)

func main() {
	e := glike.NewEngine()

	e.Use(glike.Recovery())
	e.Use(glike.DefaultLogger())

	e.GET("/", func(ctx *glike.Context) {

		ctx.String(200, "hello world")
	})

	e.GET("/hello", func(ctx *glike.Context) {
		// methods to get parameters not been implemented yet

		ctx.JSON(200, glike.H{
			"hello": "world",
			"hi":    "gin",
		})
	})

	e.GET("/string", func(ctx *glike.Context) {
		ctx.String(200, "some string")
	})

	e.Run("")
}
