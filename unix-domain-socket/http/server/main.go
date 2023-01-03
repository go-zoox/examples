package main

import (
	"fmt"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/zoox"
	zd "github.com/go-zoox/zoox/default"
)

func main() {
	r := zd.Default()

	r.Static("/assets", fs.CurrentDir())

	r.Get("/", func(ctx *zoox.Context) {
		ctx.Write([]byte("helloworld"))
	})

	if err := r.Run("unix:///tmp/127.0.0.1.sock"); err != nil {
		fmt.Println("run error: ", err)
	}
}
