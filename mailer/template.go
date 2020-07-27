package mailer

import (
	"github.com/badoux/checkmail"
	"github.com/juliankoehn/mchurl/config"
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

const defaultConfirmationMail = `<h2>Welcome</h2>
<p>Your need User has been created.:</p>
<p>Sign in By {{ .Email }} and pass: {{ .Password }}</p>`

// ValidateEmail returns nil if the email is valid,
// otherwise an error indicating the reason it is invalid
func (m TemplateMailer) ValidateEmail(email string) error {
	return checkmail.ValidateFormat(email)
}

func (m *TemplateMailer) ConfirmationMail(email, password string) error {

	data := map[string]interface{}{
		"Email":    email,
		"Password": password,
	}
	return m.Mailer.Mail(
		email,
		string(withDefault(m.Config.Mailer.Subjects.Confirmation, "Account created")),
		enforceRelativeURL(m.Config.Mailer.Templates.Confirmation),
		defaultConfirmationMail,
		data,
	)
}
