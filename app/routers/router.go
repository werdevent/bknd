package routers

import (
	"net/http"

	"github.com/GeorgeHN666/werdevent-backend/app/handlers"
	"github.com/go-chi/chi/v5"
)

func HandleRoutes() http.Handler {
	mux := chi.NewRouter()
	// Login section
	mux.Post("/log", handlers.LoginUser)
	// Password Recovery section
	mux.Get("/recover", handlers.StartRecoveryProcess)
	mux.Get("/verifier", handlers.ValidateRecoveryCode)
	mux.Post("/changer", handlers.ChangePassword)

	//Sign up process
	mux.Post("/user", handlers.CreateUser)
	mux.Get("/email_verifier", handlers.VerifyEmail)
	mux.Post("/update_details_user", handlers.UpdateUserDetails)
	mux.Post("/update_loc", handlers.UpdateLocation)
	mux.Post("/update_details", handlers.UpdateConsumerDetails)

	return mux
}
