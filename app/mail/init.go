package mail

import (
	"os"
	"strconv"

	"github.com/GeorgeHN666/werdevent-backend/app/models"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
}

func StartMailing(mailType string) *Config {

	cfg := &Config{}

	if mailType == models.SENDER_MAILER {

		port, _ := strconv.Atoi(os.Getenv(models.MAIL_SENDER_PORT))
		cfg.Host = os.Getenv(models.MAIL_SENDER_HOST)
		cfg.User = os.Getenv(models.MAIL_SENDER_USER)
		cfg.Port = port
		cfg.Password = os.Getenv(models.MAIL_SENDER_PASSWORD)
	} else {
		port, _ := strconv.Atoi(os.Getenv(models.MAIL_RECEIVER_PORT))
		cfg.Host = os.Getenv(models.MAIL_RECEIVER_HOST)
		cfg.User = os.Getenv(models.MAIL_RECEIVER_USER)
		cfg.Port = port
		cfg.Password = os.Getenv(models.MAIL_RECEIVER_PASSWORD)
	}

	return cfg
}
