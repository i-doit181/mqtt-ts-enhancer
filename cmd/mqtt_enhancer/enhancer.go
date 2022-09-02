package mqtt_enhancer

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "mqtt_enhancer",
	Version: version,
	Short:   "mqtt_enhancer -- a simple CLI to enhance json-based mqtt messages",
	Long:    `mqtt_enhancer enhanced json-based mqtt-messages with a timestamp`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Error("There was an error")
		os.Exit(1)
	}
}
