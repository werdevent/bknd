package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRoutes() http.Handler {

	mux := chi.NewRouter()

	return mux
}
