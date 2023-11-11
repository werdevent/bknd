package mail

import (
	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/app/utils"
	"gopkg.in/gomail.v2"
)

func (c *Config) SendVerificationCode(payload *models.SendMailPayload) error {

	HTML, err := utils.DownloadMailTemplate(payload.Template)
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetBody("text/html", string(HTML))
	mailer.SetHeader("To", payload.Target)
	mailer.SetHeader("From", c.User)
	mailer.SetHeader("subject", payload.Subject)

	dialer := gomail.NewDialer(c.Host, c.Port, c.User, c.Password)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
