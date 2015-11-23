package cmd

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/mannkind/wink-mqtt/controller"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version string = "0.1.0"

var cfgFile string
var reload = make(chan bool)

// WinkCmd - The root Wink commands
var WinkCmd = &cobra.Command{
	Use:   "wink",
	Short: "A replace 'firmware' for Wink",
	Long:  "A replace 'firmware' for Wink",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			controller := controller.WinkController{}

			if err := viper.Unmarshal(&controller); err != nil {
				log.Panicf("Error unmarshaling configuration: %s", err)
			}

			if err := controller.Start(); err != nil {
				log.Panicf("Error starting handlers.Controller: %s", err)
			}

			select {
			case <-reload:
				if controller.Client != nil && controller.Client.IsConnected() {
					controller.Client.Disconnect(0)
					time.Sleep(500 * time.Millisecond)
				}
				continue
			}
		}
	},
}

// Execute - Adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := WinkCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.Printf("Wink Version: %s", version)

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

	WinkCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".wink.yaml", "The path to the configuration file")
}
