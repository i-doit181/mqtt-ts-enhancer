package mqtt

import (
	"encoding/base64"
	"encoding/json"
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
	}).Debug("Received message")
	var payloadMap map[string]interface{}
	err := json.Unmarshal(msg.Payload(), &payloadMap)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"message":   msg.Payload(),
			"messageID": msg.MessageID(),
			"topic":     msg.Topic(),
		}).Fatalf("Message cannot be parsed!")
	}
	payloadMap["timestamp"] = time.Now().Unix()
	newMessage, _ := json.Marshal(payloadMap)
	publish(client, msg.Topic(), &newMessage)
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
	opts.SetKeepAlive(5 * time.Second)
	opts.OnConnectionLost = connectLostHandler
	opts.OnConnect = connectHandler
	opts.OnReconnecting = reconnectHandler
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(false)
	//
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return &client, token.Error()
	}

	return &client, nil
}

func publish(client mqtt.Client, topic string, message *[]byte) {
	token := client.Publish(topic+"_extended", 1, true, *message)
	token.Wait()
	time.Sleep(time.Second)
}

func Sub(client mqtt.Client, topic *string) {
	token := client.Subscribe(*topic, 1, nil)
	token.Wait()
	log.WithFields(log.Fields{
		"topic": *topic,
	}).Info("Subscribed to topic")
}
