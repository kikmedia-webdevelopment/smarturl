package mailer

import (
	"net/url"
	"regexp"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/models"
)

// Mailer defines the interface a mailer must implement.
type Mailer interface {
	RecoveryMail(user *models.User, referrerURL string) error
	EmailChangeMail(user *models.User, referrerURL string) error
	ValidateEmail(email string) error
}

// NewMailer returns a new mailer
func NewMailer(instanceConfig *config.Configuration) Mailer {
	if instanceConfig.SMTP.Host == "" {
		return &noopMailer{}
	}

	return &TemplateMailer{
		SiteURL: instanceConfig.Web.BaseURL,
		Config:  instanceConfig,
		Mailer: &MailmeMailer{
			Host:    instanceConfig.SMTP.Host,
			Port:    instanceConfig.SMTP.Port,
			User:    instanceConfig.SMTP.User,
			Pass:    instanceConfig.SMTP.Pass,
			From:    instanceConfig.SMTP.AdminEmail,
			BaseURL: instanceConfig.Web.BaseURL,
		},
	}
}

func withDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func getSiteURL(referrerURL, siteURL, filepath, fragment string) (string, error) {
	baseURL := siteURL
	if filepath == "" && referrerURL != "" {
		baseURL = referrerURL
	}

	site, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	if filepath != "" {
		path, err := url.Parse(filepath)
		if err != nil {
			return "", err
		}
		site = site.ResolveReference(path)
	}
	site.Fragment = fragment
	return site.String(), nil
}

var urlRegexp = regexp.MustCompile(`^https?://[^/]+`)

func enforceRelativeURL(url string) string {
	return urlRegexp.ReplaceAllString(url, "")
}
