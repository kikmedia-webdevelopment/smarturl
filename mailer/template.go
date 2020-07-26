package mailer

import (
	"github.com/badoux/checkmail"
	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/models"
)

// TemplateMailer will send mail and use templates from the site for easy mail styling
type TemplateMailer struct {
	SiteURL string
	Config  *config.Configuration
	Mailer  *MailmeMailer
}

const defaultRecoveryMail = `<h2>Reset password</h2>
<p>Follow this link to reset the password for your user:</p>
<p><a href="{{ .ConfirmationURL }}">Reset password</a></p>`

const defaultEmailChangeMail = `<h2>Confirm email address change</h2>
<p>Follow this link to confirm the update of your email address from {{ .Email }} to {{ .NewEmail }}:</p>
<p><a href="{{ .ConfirmationURL }}">Change email address</a></p>`

// ValidateEmail returns nil if the email is valid,
// otherwise an error indicating the reason it is invalid
func (m TemplateMailer) ValidateEmail(email string) error {
	return checkmail.ValidateFormat(email)
}

// EmailChangeMail sends an email change confirmation mail to a user
func (m *TemplateMailer) EmailChangeMail(user *models.User, referrerURL string) error {
	url, err := getSiteURL(referrerURL, m.Config.Web.BaseURL, m.Config.Mailer.URLPaths.EmailChange, "email_change_token="+user.EmailChangeToken)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"SiteURL":         m.Config.Web.BaseURL,
		"ConfirmationURL": url,
		"Email":           user.Email,
		"NewEmail":        user.EmailChange,
		"Token":           user.EmailChangeToken,
	}

	return m.Mailer.Mail(
		user.EmailChange,
		string(withDefault(m.Config.Mailer.Subjects.EmailChange, "Confirm Email Change")),
		enforceRelativeURL(m.Config.Mailer.Templates.EmailChange),
		defaultEmailChangeMail,
		data,
	)
}

// RecoveryMail sends a password recovery mail
func (m *TemplateMailer) RecoveryMail(user *models.User, referrerURL string) error {
	url, err := getSiteURL(referrerURL, m.Config.Web.BaseURL, m.Config.Mailer.URLPaths.Recovery, "recovery_token="+user.RecoveryToken)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"SiteURL":         m.Config.Web.BaseURL,
		"ConfirmationURL": url,
		"Email":           user.Email,
		"Token":           user.RecoveryToken,
	}

	return m.Mailer.Mail(
		user.Email,
		string(withDefault(m.Config.Mailer.Subjects.Recovery, "Reset Your Password")),
		enforceRelativeURL(m.Config.Mailer.Templates.Recovery),
		defaultRecoveryMail,
		data,
	)
}

// Send can be used to send one-off emails to users
func (m TemplateMailer) Send(user *models.User, subject, body string, data map[string]interface{}) error {
	return m.Mailer.Mail(
		user.Email,
		subject,
		"",
		body,
		data,
	)
}
