package encoders

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ReadJSON will read incoming json incoming from the body and it will unmarshall it in to a given structure
func ReadJSON(r *http.Request, data interface{}) error {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("bi dimensional structures are not allowed")
	}

	return nil
}

// WriteJSON is short hand for sending json to the client
func WriteJSON(w http.ResponseWriter, data interface{}, status int) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}
