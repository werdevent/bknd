package handlers

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/GeorgeHN666/werdevent-backend/app/DB"
	"github.com/GeorgeHN666/werdevent-backend/app/encoders"
	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/app/utils"
	"github.com/GeorgeHN666/werdevent-backend/constants"
)

// UpdateConsumerDetails Update the details of the customer of the given user id
func UpdateConsumerDetails(w http.ResponseWriter, r *http.Request) {

	var details models.ConsumerDetails
	uuid := r.URL.Query().Get(constants.HEADER_UUID)

	err := encoders.ReadJSON(r, &details)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.PAYLOAD_ERROR), http.StatusNotAcceptable)
		return
	}

	filteredModel := utils.FilterEmptyStructsWithSuffix(&details, "consumer_details")

	err = db.StartDatabase(os.Getenv("DB"), constants.DATABASE_NAME).UpdateConsumerDetails(uuid, filteredModel)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v - %v", err.Error(), constants.INTERNAL_ERROR), http.StatusInternalServerError)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	res.Error = false
	res.Message = fmt.Sprintf("Details successfuly updated - %v", constants.SUCCESSFUL)

	encoders.WriteJSON(w, res, http.StatusOK)

}
