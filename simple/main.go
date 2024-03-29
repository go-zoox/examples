package main

import (
	"github.com/go-zoox/fs"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

func main() {
	r := defaults.Application()

	r.Static("/assets", fs.CurrentDir())

	r.Get("/", func(ctx *zoox.Context) {
		ctx.Write([]byte("helloworld7"))
	})

	r.Get("/panic", func(ctx *zoox.Context) {
		var a []int
		a[0] = 1
	})

	r.Post("/body", func(ctx *zoox.Context) {
		// ctx.JSON(200, ctx.Bodies())

		type Body struct {
			A int64  `json:"a"`
			B int64  `json:"b"`
			C string `json:"c"`
		}

		var body Body
		if err := ctx.BindJSON(&body); err != nil {
			panic(err)
		}

		ctx.JSON(200, body)
	})

	v1 := r.Group("/v1")
	{
		v1.Get("/", func(ctx *zoox.Context) {
			ctx.Write([]byte("v1"))
		})
		v1.Get("/hello", func(ctx *zoox.Context) {
			ctx.JSON(200, zoox.H{
				"hello": "world",
			})
		})
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
