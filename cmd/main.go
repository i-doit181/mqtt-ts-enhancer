package main

import (
	"github.com/shelly-ts-enhancer/cmd/mqtt_enhancer"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Run mqtt_enhancer CLI!")
	mqtt_enhancer.Execute()
}
