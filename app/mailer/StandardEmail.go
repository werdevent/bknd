package mailer

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendStandardEmail(targetEmail, subject string, data map[string]interface{}, templatePath string) error {

	mail := gomail.NewMessage()

	mail.SetHeader("From", os.Getenv("sender-user"))
	mail.SetHeader("To", targetEmail)
	mail.SetHeader("Subject", subject)

	content, err := EmailTemplates.ReadFile(templatePath)
	if err != nil {
		return err
	}
	HTML, err := EmbedDataToEmail(string(content), data)
	if err != nil {
		return err
	}
	mail.SetBody("text/html", HTML)
	_, err = strconv.Atoi(os.Getenv("sender-port"))
	if err != nil {
		return err
	}
	send := gomail.NewDialer("mx106.hostgator.mx", 465, "contact@zkaia.com", "log.Fatal(1$)")

	err = send.DialAndSend(mail)
	if err != nil {
		return err
	}
	return nil
}
