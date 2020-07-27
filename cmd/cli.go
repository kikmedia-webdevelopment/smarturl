package cmd

import (
	"net/url"
	"os"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/models"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	uri  string
	code string
)

func createUser(config *config.Configuration, cmd *cobra.Command) {
	var email string
	emailFlag := cmd.Flags().Lookup("email")
	if emailFlag != nil {
		if emailFlag.Value.String() != "" {
			email = emailFlag.Value.String()
		}
	}

	db, err := storage.Dial(config)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()
	autoMigrate(db)

	user, password, err := models.NewUser(db, config, email, "")
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	logrus.Infof("User `%s` with password `%s` created", user.Email, password)
}

func delete(config *config.Configuration, cmd *cobra.Command) {
	var code string
	codeFlag := cmd.Flags().Lookup("code")
	if codeFlag != nil {
		if codeFlag.Value.String() != "" {
			code = codeFlag.Value.String()
		}
	}

	if code == "" {
		logrus.Error("Code is missing")
		os.Exit(1)
	}

	db, err := storage.Dial(config)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()
	autoMigrate(db)

	err = models.DeleteEntry(db, code)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	logrus.Infof("ShortURL by %s has been deleted", code)
}

func create(config *config.Configuration, cmd *cobra.Command) {
	var code string
	urlFlag := cmd.Flags().Lookup("url")

	if urlFlag == nil {
		logrus.Error("URL is missing")
		os.Exit(1)
	}

	if urlFlag.Value.String() == "" {
		logrus.Error("URL is missing")
		os.Exit(1)
	}

	uri := urlFlag.Value.String()
	if _, err := url.ParseRequestURI(uri); err != nil {
		logrus.Errorf("Error parsing URL: %+v", err)
		os.Exit(1)
	}

	codeFlag := cmd.Flags().Lookup("code")
	if codeFlag != nil {
		if codeFlag.Value.String() != "" {
			code = codeFlag.Value.String()
		}
	}

	db, err := storage.Dial(config)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()
	autoMigrate(db)

	_, err = models.CreateEntry(db, config.DB.IDLength, uri, code)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	logrus.Infof("ShortURL by %s has been created", code)
}
