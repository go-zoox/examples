package main

import (
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/logger"
)

func main() {
	response, err := fetch.Get("http://127.0.0.1", &fetch.Config{
		UnixDomainSocket: "/tmp/127.0.0.1.sock",
	})
	if err != nil {
		logger.Errorf("request error: %v", err)
		return
	}

	logger.Infof("response: %s", response.String())
}
