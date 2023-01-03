package main

import (
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/logger"
)

func main() {
	response, err := fetch.Get("https://zsxxx.com", &fetch.Config{
		TLSCertificateFile: fs.JoinPath(fs.CurrentDir(), "../server/server.crt"),
		UnixDomainSocket:   "/tmp/zsxxx.com.sock",
	})
	if err != nil {
		logger.Errorf("request error: %v", err)
		return
	}

	logger.Infof("response: %s", response.String())
}
