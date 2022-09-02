package mqtt_enhancer

import (
	mqtt "github.com/shelly-ts-enhancer/internal/mqtt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var mqttBroker string
var topic string

var convertCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"cvt"},
	Short:   "Connect and enhance payload",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log.WithFields(log.Fields{
			"mqtt":  mqttBroker,
			"topic": topic,
		}).Info("Here the magic happens ")
		client, err := mqtt.Connect(&mqttBroker)
		if err != nil {
			log.WithError(err).Error("Something went wrong!")
		} else {
			mqtt.Sub(*client, &topic)
		}
	},
}

func init() {
	convertCmd.Flags().StringVarP(&mqttBroker, "mqtt", "m", "", "MQTT Broker")
	convertCmd.Flags().StringVarP(&topic, "topic", "t", "", "Topic")
	rootCmd.AddCommand(convertCmd)
}
