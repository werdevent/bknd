package models

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
