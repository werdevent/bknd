package models

type StandardCredential struct {
	Signature Signature `json:"signature"`
	Payload   string    `json:"payload"`
}

type AccessCredentials struct {
	Signature Signature `json:"signature"`
	Role      int64     `json:"role"`
	Logged    bool      `json:"logged"`
}

type InternalPaylod struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	UI    string `json:"ui"`
	BI    string `json:"bi"`
	CI    string `json:"CI"`
}
