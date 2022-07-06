package main

import (
	"github.com/go-zoox/jsonrpc"
	"github.com/go-zoox/uuid"
	"github.com/go-zoox/zoox"
	zd "github.com/go-zoox/zoox/default"
)

func main() {
	r := zd.Default()

	rpc := jsonrpc.NewServer()

	rpc.Register("uuid", func(ctx *jsonrpc.Context, params jsonrpc.Params) jsonrpc.Result {
		return jsonrpc.Result{
			"uuid": uuid.V4(),
		}
	})

	rpc.Register("ip", func(ctx *jsonrpc.Context, params jsonrpc.Params) jsonrpc.Result {
		return jsonrpc.Result{
			"ip": ctx.Request.RemoteAddr,
		}
	})

	r.Post("/jsonrpc", zoox.WrapH(rpc))

	r.Run(":8009")
}
