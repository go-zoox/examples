package main

import (
	"github.com/go-zoox/fs"
	"github.com/go-zoox/zoox"
	zd "github.com/go-zoox/zoox/default"
)

func main() {
	r := zd.Default()

	r.TLSKeyFile = "./server.key"
	r.TLSCertFile = "./server.crt"

	r.Static("/assets", fs.CurrentDir())

	r.Get("/", func(ctx *zoox.Context) {
		ctx.Write([]byte("helloworld"))
	})

	r.Run(":9996")
}
