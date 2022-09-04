package mqtt_enhancer

import (
	"github.com/matryer/is"
	"testing"
)

func Test_StartWithNotAvailableBroker(t *testing.T) {
	isErr := is.New(t)
	rootCmd.SetArgs([]string{"start", "-m", "nobroker:1883", "-t", "test"})
	err := Execute()
	isErr.True(err != nil)
}
