package api

import (
	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/mailer"
	"github.com/juliankoehn/mchurl/models"
)

// Mailer returns a new mailer instance
func Mailer(config *config.Configuration) mailer.Mailer {
	return mailer.NewMailer(config)
}

func sendConfirmation(u *models.User, password string, mailer mailer.Mailer) error {
	return mailer.ConfirmationMail(u.Email, password)
}
