package main

import (
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/logger"
)

// curl --cacert $PWD/server.crt  https://zsxxx.com:9996
func main() {
	// if no https certificate set:
	//		error:  ERROR request error: ErrSendingRequest(3):  error sending request, err: Get "https://zsxxx.com:9996": x509: “*.zsxxx.com” certificate is not trusted(Please check your network, maybe use bad proxy or network offline)
	//
	response, err := fetch.Get("https://zsxxx.com:9996", &fetch.Config{
		TLSCertificateFile: fs.JoinPath(fs.CurrentDir(), "../server.crt"),
	})
	if err != nil {
		logger.Errorf("request error: %v", err)
		return
	}

	logger.Infof("response: %s", response.String())
}
