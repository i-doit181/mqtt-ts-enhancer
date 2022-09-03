package mqtt_enhancer

import (
	mqtt "github.com/shelly-ts-enhancer/internal/mqtt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var mqttBroker string
var topics string

var convertCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"cvt"},
	Short:   "Connect and enhance payload",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		keepAlive := make(chan os.Signal)
		signal.Notify(keepAlive, os.Interrupt, syscall.SIGTERM)

		log.WithFields(log.Fields{
			"mqtt":   mqttBroker,
			"topics": topics,
		}).Info("Enhance timestamp to messages of following topics")

		client, err := mqtt.Connect(&mqttBroker)

		if err != nil {
			log.WithError(err).Fatal("Something went wrong!")
		} else {
			topicList := strings.SplitN(topics, ",", -1)
			for _, topic := range topicList {
				mqtt.Sub(*client, &topic)
			}
		}

		<-keepAlive
	},
}

func init() {
	convertCmd.Flags().StringVarP(&mqttBroker, "mqtt", "m", "", "MQTT Broker")
	convertCmd.Flags().StringVarP(&topics, "topics", "t", "", "Comma-driven topic List")
	rootCmd.AddCommand(convertCmd)
}
