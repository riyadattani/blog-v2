package httpserver

import (
	"time"
)

type ServerConfig struct {
	Port             string        `envconfig:"PORT" required:"true" default:"8080"`
	HTTPReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" required:"true" default:"2s"`
	HTTPWriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" required:"true" default:"2s"`
}

func (c ServerConfig) TCPAddress() string {
	return ":" + c.Port
}
