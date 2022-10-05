package main

import (
	"github.com/go-zoox/logger"
	"github.com/go-zoox/zoox"
	zd "github.com/go-zoox/zoox/default"
)

func main() {
	socket := "unix:///tmp/zoox.sock"
	app := zd.Default()

	app.Get("/", func(ctx *zoox.Context) {
		ctx.String(200, "Hello, world!")
	})

	if err := app.Run(socket); err != nil {
		logger.Error("Failed to run application: %v", err)
	}
}

// test with curl:
//	curl --unix-socket /tmp/zoox.sock http://localhost/
// references:
//   - https://cloud.tencent.com/developer/article/1786322
//   - https://www.cnblogs.com/walkinginthesun/p/10397539.html
//   - https://gist.github.com/teknoraver/5ffacb8757330715bcbcc90e6d46ac74
