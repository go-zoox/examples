package main

import (
	"time"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

func main() {
	r := defaults.Application()

	r.Static("/assets", fs.CurrentDir())

	r.Get("/", func(ctx *zoox.Context) {
		ctx.Write([]byte("helloworld"))
	})

	r.Get("/sse", func(ctx *zoox.Context) {
		// ctx.Status(200)
		// ctx.SetHeader("Content-Type", "text/event-stream")
		// ctx.SetHeader("Cache-Control", "no-cache")
		// ctx.SetHeader("Connection", "keep-alive")
		// ctx.SetHeader("Access-Control-Allow-Origin", "*")

		// ctx.Write([]byte("retry: 10000\n"))
		// ctx.Write([]byte("event: connecttime\n"))
		// ctx.Writer.Flush()

		// i := 0
		// for {
		// 	if i > 10 {
		// 		break
		// 	}

		// 	ctx.Write([]byte(fmt.Sprintf("id: %d\n", i)))
		// 	ctx.Write([]byte(fmt.Sprintf("data: %s\n\n", time.Now())))
		// 	ctx.Writer.Flush()

		// 	i += 1
		// 	time.Sleep(1 * time.Second)
		// }

		sse := ctx.SSE()
		sse.Retry(10000 * time.Second)
		// sse.Event("connecttime")

		i := 0
		for {
			if i > 10 {
				break
			}

			sse.Event("connecttime", time.Now().String())

			i += 1
			time.Sleep(1 * time.Second)
		}
	})

	r.Run(":8080")
}
