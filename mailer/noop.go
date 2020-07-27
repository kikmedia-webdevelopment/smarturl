package mailer

type noopMailer struct {
}

func (m noopMailer) ConfirmationMail(email, password string) error {
	return nil
}
