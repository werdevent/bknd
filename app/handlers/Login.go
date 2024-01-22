package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	db "github.com/GeorgeHN666/werdevent-backend/app/DB"
	"github.com/GeorgeHN666/werdevent-backend/app/encoders"
	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/app/utils"
	"github.com/GeorgeHN666/werdevent-backend/constants"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var u models.User
	err := encoders.ReadJSON(r, &u)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.PAYLOAD_ERROR)), http.StatusNotAcceptable)
		return
	}

	// check user existance
	user, err := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(u.Email)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR)), http.StatusInternalServerError)
		return
	}

	// check pwd
	match, err := encoders.CompareHashedPWD(user.Password, u.Password)
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.BAD_CREDENTIALS)), http.StatusForbidden)
		return
	}

	if !match {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("Incorrect data - %v", constants.INCORRECT_DATA)), http.StatusForbidden)
		return
	}

	// send required models

	data := models.InternalPaylod{
		Name:  user.Name,
		Email: user.Email,
	}

	sign, err := json.Marshal(models.Signature{
		Sign:    models.SIGNATURE,
		Type:    models.SIGN_TYPE,
		Created: fmt.Sprintf("%v", time.Now().UnixNano()),
		UI:      user.ID.Hex(),
		BI:      user.Consumer.Buyer_ID,
		CI:      user.Consumer.Creator_ID,
	})
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR)), http.StatusInternalServerError)
		return
	}

	encryptedData, err := encoders.EncryptPayload(string(sign), []byte(models.PKEY))
	if err != nil {
		http.Error(w, utils.ThrowJSONerror(fmt.Sprintf("Internal error - %v", constants.INTERNAL_ERROR)), http.StatusInternalServerError)
		return
	}

	payload := &models.StandardCredential{
		Signature: encryptedData,
		Payload:   data,
		Role:      int64(user.Role),
		Logged:    true,
	}

	var res struct {
		Error   bool                       `json:"error"`
		Message string                     `json:"message"`
		Payload *models.StandardCredential `json:"payload,omitempty"`
	}
	res.Error = false
	res.Message = fmt.Sprintf("Successfully login - %v", constants.SUCCESSFUL)
	res.Payload = payload

	encoders.WriteJSON(w, res, http.StatusOK)
}
