package mqtt

import (
	"encoding/base64"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.WithFields(log.Fields{
		"topic":     msg.Topic(),
		"messageID": msg.MessageID(),
	}).Info("Received message\n")
	//here validate json and add timestamp. Finally publish to other topic
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info("Connected!")
}

var reconnectHandler mqtt.ReconnectHandler = func(client mqtt.Client, options *mqtt.ClientOptions) {
	log.WithFields(log.Fields{
		"retry": options.ConnectRetry,
	}).Info("Reconnected!")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.WithError(err).Error("Connect lost!")
}

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")

func Connect(broker *string) (*mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID("mqtt_enhancer_" + escaper.Replace(base64.RawURLEncoding.EncodeToString([]byte(uuid.NewString()))))
	opts.OnConnectionLost = connectLostHandler
	opts.OnConnect = connectHandler
	opts.OnReconnecting = reconnectHandler
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetAutoReconnect(true)
	//
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return &client, token.Error()
	}

	return &client, nil
}

func Publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func Sub(client mqtt.Client, topic *string) {
	token := client.Subscribe(*topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", *topic)
}
