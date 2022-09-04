package mqtt_enhancer

import (
	"testing"

	"github.com/matryer/is"
)

func Test_ExecuteCommandWithWrongArgs(t *testing.T) {
	isErr := is.New(t)
	rootCmd.SetArgs([]string{"dummy"})
	err := Execute()
	isErr.True(err != nil)
}

func Test_RootCmdWithoutArgs(t *testing.T) {
	isNoErr := is.New(t)
	rootCmd.SetArgs([]string{})
	err := Execute()
	isNoErr.NoErr(err)
}

func Test_SubCmdWithHelp(t *testing.T) {
	isNoErr := is.New(t)
	rootCmd.SetArgs([]string{"start", "--help"})
	err := Execute()
	isNoErr.NoErr(err)
}
