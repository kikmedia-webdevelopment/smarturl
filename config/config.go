package config

import (
	"errors"

	"github.com/spf13/viper"
)

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
	Secret string           `yaml:"Secret"`
	Web    WebConfiguration `yaml:"Web"`
	DB     DBConfiguration  `yaml:"DB"`
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
