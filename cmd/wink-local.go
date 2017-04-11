package cmd

import (
	"log"

	"github.com/mannkind/wink-local/controller"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// WinkLocal - The root Wink commands
var WinkLocal = &cobra.Command{
	Use:   "wink-local",
	Short: "A local-control replacement for the Wink Hub",
	Long:  "A local-control replacement for the Wink Hub",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.WinkController{}

		if err := viper.Unmarshal(&controller); err != nil {
			log.Panicf("Error unmarshaling configuration: %s", err)
		}

		if err := controller.Start(); err != nil {
			log.Panicf("Error starting handlers.Controller: %s", err)
		}

		select {}
	},
}

// Execute - Adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := WinkLocal.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfgFile)
		log.Printf("Loading Configuration %s", cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error Loading Configuration: %s ", err)
		}
		log.Printf("Loaded Configuration %s", cfgFile)
	})

	WinkLocal.PersistentFlags().StringVarP(&cfgFile, "config", "c", "wink-local.yaml", "The path to the configuration file")
}
