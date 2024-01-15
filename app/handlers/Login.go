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
	"github.com/GeorgeHN666/werdevent-backend/constants"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var u models.User
	err := encoders.ReadJSON(r, &u)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.PAYLOAD_ERROR), http.StatusNotAcceptable)
		return
	}

	// check user existance
	user, err := db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).GetUser(u.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.Internal_DB_ERROR), http.StatusInternalServerError)
		return
	}

	// check pwd
	match, err := encoders.CompareHashedPWD(user.Password, u.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR), http.StatusForbidden)
		return
	}

	if !match {
		http.Error(w, fmt.Sprintf("Incorrect data - %v", constants.INCORRECT_DATA), http.StatusForbidden)
		return
	}

	// send required models

	data, err := json.Marshal(models.InternalPaylod{
		Name:  user.Name,
		Email: user.Email,
		UI:    user.ID.Hex(),
		BI:    user.Consumer.Buyer_ID,
		CI:    user.Consumer.Creator_ID,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error - %v", constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}

	encryptedData, err := encoders.EncryptPayload(string(data), []byte(models.PKEY))
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error - %v", constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}

	sign := models.Signature{
		Sign:    models.SIGNATURE,
		Type:    models.SIGN_TYPE,
		Created: fmt.Sprintf("%v", time.Now().UnixNano()),
	}

	payload := &models.StandardCredential{
		Signature: sign,
		Payload:   encryptedData,
	}

	Access := &models.AccessCredentials{
		Signature: sign,
		Role:      int64(u.Role),
	}

	var res struct {
		Error   bool                       `json:"error"`
		Message string                     `json:"message"`
		Payload *models.StandardCredential `json:"payload"`
		Access  *models.AccessCredentials  `json:"access"`
	}
	res.Error = false
	res.Message = fmt.Sprintf("Successfully login - %v", constants.SUCCESSFUL)
	res.Payload = payload
	res.Access = Access

	encoders.WriteJSON(w, res, http.StatusOK)
}
