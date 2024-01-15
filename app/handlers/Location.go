package handlers

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/GeorgeHN666/werdevent-backend/app/DB"
	"github.com/GeorgeHN666/werdevent-backend/app/encoders"
	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/constants"
)

// UpdateLocation Update the default location the user is using
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var Location models.Location
	uid := r.URL.Query().Get(constants.HEADER_UUID)

	err := encoders.ReadJSON(r, &Location)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.PAYLOAD_ERROR), http.StatusNotAcceptable)
		return
	}

	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateLocation(uid, &Location)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Location successfully updated - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)

}
