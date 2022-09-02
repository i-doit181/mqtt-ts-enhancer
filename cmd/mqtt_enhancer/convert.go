package mqtt_enhancer

import (
	mqtt "github.com/shelly-ts-enhancer/internal/mqtt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var mqttBroker string
var topicOngoing string
var topicOutgoing string

var convertCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"cvt"},
	Short:   "Connect and enhance payload",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log.WithFields(log.Fields{
			"mqtt":          mqttBroker,
			"topicOngoing":  topicOngoing,
			"topicOutgoing": topicOutgoing,
		}).Info("Here the magic happens ")
		client, err := mqtt.Connect(&mqttBroker)
		if err != nil {
			log.WithError(err).Error("Something went wrong!")
		} else {
			mqtt.Sub(*client)
		}
	},
}

func init() {
	convertCmd.Flags().StringVarP(&mqttBroker, "mqtt", "m", "", "MQTT Broker")
	convertCmd.Flags().StringVarP(&topicOngoing, "consume", "c", "", "Topic where messages needs to be enhanced")
	convertCmd.Flags().StringVarP(&topicOutgoing, "produce", "p", "", "Topic where enhanced messages are published")
	rootCmd.AddCommand(convertCmd)
}
