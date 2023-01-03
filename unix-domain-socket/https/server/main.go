package main

import (
	"fmt"

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

	// r.Run(":9996")
	// r.Run()
	if err := r.Run("unix:///tmp/zsxxx.com.sock"); err != nil {
		fmt.Println("run error: ", err)
	}
}
