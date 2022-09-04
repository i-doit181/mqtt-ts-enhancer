package mqtt_enhancer

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "mqtt_enhancer",
	Version: version,
	Short:   "mqtt_enhancer -- a simple CLI to enhance json-based mqtt messages",
	Long:    `mqtt_enhancer enhanced json-based mqtt-messages with a timestamp`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Error("There was an error")
		return err
	}

	return nil
}
