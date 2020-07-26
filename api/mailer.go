package api

import "github.com/juliankoehn/mchurl/mailer"

func (a *API) Mailer() mailer.Mailer {
	return mailer.NewMailer(a.config)
}
