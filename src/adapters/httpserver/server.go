package httpserver

import (
	"net/http"
)

func NewWebServer(
	config ServerConfig,
	router http.Handler,
) *http.Server {
	return &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      router,
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
