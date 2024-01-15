package models

type StandardCredential struct {
	Signature Signature `json:"signature"`
	Payload   string    `json:"payload"`
}

type InternalPaylod struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	UI    string `json:"ui"`
	BI    string `json:"bi"`
	CI    string `json:"CI"`
}
