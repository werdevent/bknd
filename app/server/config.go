package server

import "net/http"

type Config struct {
	PORT    int
	Handler http.Handler
	Env     string
	Version string
}
