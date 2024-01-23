package middleware

import (
	"net/http"
)

var httpServer http.Server

func CreateHttpServer() {
	httpServer = http.Server{
		Handler: httpHandler,
	}
}

func RunHttpServer() error {
	return httpServer.Serve(httpL)
}

func StopHttpServer() {
	_ = httpServer.Close()
}
