package handlers

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/GeorgeHN666/werdevent-backend/app/DB"
	"github.com/GeorgeHN666/werdevent-backend/app/encoders"
	"github.com/GeorgeHN666/werdevent-backend/app/mailer"
	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/app/utils"
	"github.com/GeorgeHN666/werdevent-backend/constants"
)

// CreateUser Creates an new instance of a user if the user hasnt been already registered
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := encoders.ReadJSON(r, &user)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v ", err.Error(), constants.PAYLOAD_ERROR)), http.StatusConflict)
		return
	}

	userData, _ := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(user.Email)
	if userData != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("User exist - %v ", constants.ALREADY_EXIST)), http.StatusFound)
		return
	}

	code, _ := encoders.GenerateStandardCode(6)
	user.Recovery_Code = code
	user.Valid_Code = true
	user.Verified = false
	hash, err := encoders.HashPassword(user.Password)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v ", err.Error(), constants.PAYLOAD_ERROR)), http.StatusConflict)
		return
	}
	user.Password = hash
	// insert user
	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).InsertUser(&user)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR)), http.StatusInternalServerError)
		return
	}
	// send code
	data := make(map[string]interface{})
	data["Code"] = code
	err = mailer.SendStandardEmail(user.Email, "Codigo de verificacion", data, "email-templates/code.html")
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR)), http.StatusInternalServerError)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("User successfully created - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)
}

// VerifyEmail Verifies the user account
func VerifyEmail(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	email := r.URL.Query().Get(constants.PAYLOAD)

	userData, err := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(email)
	if userData == nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.NO_FOUNDED)), http.StatusNotFound)
		return
	}

	if !userData.Valid_Code {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("expired - %v", constants.EXPIRED)), http.StatusNotAcceptable)
		return
	}

	if userData.Recovery_Code != code {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("error with code - %v", constants.INCORRECT_DATA)), http.StatusNotAcceptable)
		return
	}

	//erase code and set validcode to false
	var user models.User
	user.Valid_Code = false
	user.Verified = true
	user.Recovery_Code = "?????????????"

	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateUserDetails(userData.ID.Hex(), &user)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("error while updating: %v - %v ", err.Error(), constants.INTERNAL_ERROR)), http.StatusInternalServerError)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Email successfuly verified - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)
}

// UpdateUserDetails Update a document in database taking security actions
func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {

	var user models.User
	uid := r.URL.Query().Get(constants.HEADER_UUID)

	err := encoders.ReadJSON(r, &user)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v ", err.Error(), constants.PAYLOAD_ERROR)), http.StatusConflict)
		return
	}

	user.Email = ""
	user.Password = ""
	user.Role = 0
	user.Created_At = ""

	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateUserDetails(uid, &user)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR)), http.StatusBadRequest)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Data updated - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)

}
