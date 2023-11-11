package server

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer(cfg *Config) error {

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.PORT),
		Handler:           cfg.Handler,
		ReadTimeout:       1 * time.Minute,
		IdleTimeout:       1 * time.Minute,
		ReadHeaderTimeout: 1 * time.Minute,
		WriteTimeout:      1 * time.Minute,
	}

	fmt.Printf("Server %s up and running in %s mode using PORT %d\n", cfg.Version, cfg.Env, cfg.PORT)

	return srv.ListenAndServe()
}
