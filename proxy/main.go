package main

import (
	"github.com/go-zoox/proxy"
	"github.com/go-zoox/zoox"
)

func main() {
	r := zoox.Default()

	r.Get("/", func(ctx *zoox.Context) {
		ctx.JSON(200, zoox.H{
			"hello": "world",
		})
	})

	// proxy api request
	r.Get("/api/*path", zoox.WrapH(proxy.NewSingleTarget("https://httpbin.org", &proxy.SingleTargetConfig{
		Rewrites: map[string]string{
			"^/api/(.*)": "/$1",
		},
	})))

	// proxy web socket connection
	r.Get("/ws/*path", zoox.WrapH(proxy.NewSingleTarget("https://httpbin.org", &proxy.SingleTargetConfig{
		Rewrites: map[string]string{
			"^/ws/(.*)": "/$1",
		},
	})))

	r.Run()
}
