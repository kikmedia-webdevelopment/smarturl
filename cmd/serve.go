package cmd

import (
	"github.com/juliankoehn/mchurl/api"
	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/juliankoehn/mchurl/stores"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = cobra.Command{
	Use:  "serve",
	Long: "Start the URL-Shortener WebService",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, serve)
	},
}

func serve(config *config.Configuration) {
	db, err := storage.Dial(config)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()

	store, err := stores.New(&config.DB)
	if err != nil {
		logrus.Fatal(err)
	}

	api.New(store, db, config)
}
