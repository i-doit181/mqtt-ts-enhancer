package cmd

import (
	"github.com/shelly-ts-enhancer/cmd/mqtt_enhancer"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.Info("Run mqtt_enhancer CLI!")
	err := mqtt_enhancer.Execute()

	if err != nil {
		os.Exit(1)
	}
}
