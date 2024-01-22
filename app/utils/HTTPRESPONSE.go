package utils

import (
	"encoding/json"

	"github.com/GeorgeHN666/werdevent-backend/app/models"
)

func ThrowJSONerror(msg string) string {

	var res models.Response
	res.Error = true
	res.Message = msg

	r, _ := json.Marshal(res)
	return string(r)
}
