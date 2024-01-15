package models

type StandardCredential struct {
	Signature string         `json:"signature"`
	Payload   InternalPaylod `json:"payload"`
	Role      int64          `json:"role"`
	Logged    bool           `json:"logged"`
}

type InternalPaylod struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
