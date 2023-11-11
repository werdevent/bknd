package models

type SendMailPayload struct {
	Target   string
	Name     string
	Subject  string
	Template string
	Content  map[string]interface{}
}
