package mailer

import (
	"embed"
	"html/template"
	"strings"
)

//go:embed email-templates
var EmailTemplates embed.FS

var EMAILDATA map[string]interface{}

func EmbedDataToEmail(emailContent string, data map[string]interface{}) (string, error) {

	tmpl, err := template.New("temp").Parse(emailContent)
	if err != nil {
		return "", err
	}

	var htmlReady strings.Builder
	err = tmpl.Execute(&htmlReady, data)
	if err != nil {
		return "", err
	}

	return htmlReady.String(), nil
}
