package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

// EmailProviderConfiguration holds email related configs
type EmailProviderConfiguration struct {
	Disabled bool `json:"disabled"`
}

// EmailContentConfiguration holds the configuration for emails, both subjects and template URLs.
type EmailContentConfiguration struct {
	Recovery    string `json:"recovery"`
	EmailChange string `json:"email_change" split_words:"true"`
}

// SMTPConfiguration is the SMTP config for the Mailer
type SMTPConfiguration struct {
	MaxFrequency time.Duration `json:"max_frequency" split_words:"true"`
	Host         string        `json:"host"`
	Port         int           `json:"port,omitempty" default:"587"`
	User         string        `json:"user"`
	Pass         string        `json:"pass,omitempty"`
	AdminEmail   string        `json:"admin_email" split_words:"true"`
}

// WebConfiguration keeps information for our WebService
type WebConfiguration struct {
	ListenAddr string `yaml:"ListenAddr`
	BaseURL    string `yaml:"BaseURL"`
	Debug      bool   `yaml:"debug"`
	Redirect   string `yaml:"404Redirect"`
}

// DBConfiguration holds information about database, database driver and connection params
type DBConfiguration struct {
	Driver   string `yaml:"Driver" required:"true"`
	URL      string `yaml:"URL" required:"true"`
	IDLength int    `yaml:"IDLength" required:"true"`
}

// Configuration holds all the configuration that applies to the shortener application.
type Configuration struct {
	Secret string            `yaml:"Secret"`
	Web    WebConfiguration  `yaml:"Web"`
	DB     DBConfiguration   `yaml:"DB"`
	SMTP   SMTPConfiguration `yaml:"smtp"`
	Mailer struct {
		Subjects  EmailContentConfiguration `yaml:"subjects"`
		Templates EmailContentConfiguration `yaml:"templates"`
		URLPaths  EmailContentConfiguration `yaml:"url_paths"`
	} `yaml:"mailer"`
}

// LoadGlobal loads the configuration from file and env variables.
func LoadGlobal(filename string) (*Configuration, error) {
	if filename == "" {
		return nil, errors.New("missing config in application rootDir")
	}

	//viper.SetConfigName(filename)
	viper.SetConfigFile(filename)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // find and read the config file
	if err != nil {
		return nil, err
	}

	config := new(Configuration)
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
