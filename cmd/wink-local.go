package cmd

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/mannkind/wink-local/controller"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version string = "0.1.3"

var cfgFile string
var reload = make(chan bool)

// WinkLocal - The root Wink commands
var WinkLocal = &cobra.Command{
	Use:   "wink-local",
	Short: "A local-control replacement for the Wink Hub",
	Long:  "A local-control replacement for the Wink Hub",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			controller := controller.WinkController{}

			if err := viper.Unmarshal(&controller); err != nil {
				log.Panicf("Error unmarshaling configuration: %s", err)
			}

			if err := controller.Start(); err != nil {
				log.Panicf("Error starting handlers.Controller: %s", err)
			}

			<-reload
			if controller.Client != nil && controller.Client.IsConnected() {
				controller.Client.Disconnect(0)
				time.Sleep(500 * time.Millisecond)
			}
		}
	},
}

// Execute - Adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := WinkLocal.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.Printf("Wink Local Version: %s", version)

	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfgFile)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("Configuration Changed: %s", e.Name)
			reload <- true
		})

		log.Printf("Loading Configuration %s", cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error Loading Configuration: %s ", err)
		}
		log.Printf("Loaded Configuration %s", cfgFile)
	})

	WinkLocal.PersistentFlags().StringVarP(&cfgFile, "config", "c", "wink-local.yaml", "The path to the configuration file")
}
