package cmd

import (
	"os"

	"github.com/juliankoehn/mchurl/api"
	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/stores"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configFile = ""

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = cobra.Command{
		Use: "mchurl",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info(args)
			execWithConfig(cmd, serve)
		},
	}
	createCmd = cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			execWithConfigAndCmd(cmd, create)
		},
	}
	deleteCmd = cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			execWithConfigAndCmd(cmd, delete)
		},
	}

	// users
	createUserCmd = cobra.Command{
		Use: "createUser",
		Run: func(cmd *cobra.Command, args []string) {
			execWithConfigAndCmd(cmd, createUser)
		},
	}
)

func init() {
	rootCmd.AddCommand(&serveCmd, &createCmd, &deleteCmd, &createUserCmd)
	createCmd.Flags().StringP("url", "u", "", "the targeted URI")
	createCmd.Flags().StringP("code", "c", "", "specific short code for the URL")
	createCmd.MarkFlagRequired("url")

	deleteCmd.Flags().StringP("code", "c", "", "specific short code for the URL")
	deleteCmd.MarkFlagRequired("code")

	createUserCmd.Flags().StringP("email", "e", "", "email of the user to create")
	createUserCmd.MarkFlagRequired("email")
}

// Execute will setup and return the root command
func Execute() error {
	// rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "the config file to use")
	// rootCmd.AddCommand(&serveCmd, createCmd)

	return rootCmd.Execute()
}

func execWithConfigAndCmd(cmd *cobra.Command, fn func(config *config.Configuration, cmd *cobra.Command)) {
	if configFile == "" {
		configFile = "config.yaml"
	}
	config, err := config.LoadGlobal(configFile)
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %+v", err)
	}
	fn(config, cmd)
}

func execWithConfig(cmd *cobra.Command, fn func(config *config.Configuration)) {
	if configFile == "" {
		configFile = "config.yaml"
	}
	config, err := config.LoadGlobal(configFile)
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %+v", err)
	}
	fn(config)
}

var serveCmd = cobra.Command{
	Use:  "serve",
	Long: "Start the URL-Shortener WebService",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, serve)
	},
}

func serve(config *config.Configuration) {

	store, err := stores.New(&config.DB)
	if err != nil {
		logrus.Fatal(err)
	}

	api.New(store, config)

	_ = store

	logrus.Info(config)
	logrus.Info("Hello from serveCmd")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
