package handlers

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/GeorgeHN666/werdevent-backend/app/DB"
	"github.com/GeorgeHN666/werdevent-backend/app/encoders"
	"github.com/GeorgeHN666/werdevent-backend/app/mailer"
	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/constants"
)

// StartRecoveryProcess Will start the password recovert process
func StartRecoveryProcess(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get(constants.PAYLOAD)
	// check if user exist
	user, err := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}
	// create code
	code, err := encoders.GenerateStandardCode(6)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}
	// save code
	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateUserDetails(user.ID.Hex(), &models.User{Recovery_Code: code, Valid_Code: true})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}
	// send code through email
	data := make(map[string]interface{})
	data["Code"] = code
	err = mailer.SendStandardEmail(user.Email, "Codigo de verificacion", data, "email-templates/code.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Recovery process successfuly started - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)

}

func ValidateRecoveryCode(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get(constants.PAYLOAD)
	code := r.URL.Query().Get(constants.CODES)
	// get user
	user, err := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}
	// validete code
	if !user.Valid_Code {
		http.Error(w, fmt.Sprintf("Invalid code - %v", constants.EXPIRED), http.StatusForbidden)
		return
	}

	if user.Recovery_Code != code {
		http.Error(w, fmt.Sprintf("Incorrect data - %v", constants.INCORRECT_DATA), http.StatusNotAcceptable)
		return
	}

	// generate another code
	Authcode, err := encoders.GenerateAlphanumericCode(8)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}

	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateUserDetails(user.ID.Hex(), &models.User{Recovery_Code: Authcode, Valid_Code: true})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}

	// send autorization code
	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Auth    string `json:"werdau"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Successfully verified - %v", constants.SUCCESSFUL)
	res.Auth = Authcode

	encoders.WriteJSON(w, res, http.StatusOK)

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get(constants.PAYLOAD)
	code := r.URL.Query().Get(constants.CODES)

	var u models.User
	err := encoders.ReadJSON(r, &u)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INCORRECT_DATA), http.StatusNotAcceptable)
		return
	}

	// get user
	user, err := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}
	// validate code
	if !user.Valid_Code {
		http.Error(w, fmt.Sprintf("Invalid code - %v", constants.EXPIRED), http.StatusForbidden)
		return
	}

	if user.Recovery_Code != code {
		http.Error(w, fmt.Sprintf("Incorrect data - %v", constants.INCORRECT_DATA), http.StatusNotAcceptable)
		return
	}

	// encrypt pwd
	pwd, err := encoders.HashPassword(u.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INCORRECT_DATA), http.StatusNotAcceptable)
		return
	}

	// save pwd
	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateUserDetails(user.ID.Hex(), &models.User{Password: pwd, Valid_Code: false, Recovery_Code: "?????????????"})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Password successfully changed - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)

}
