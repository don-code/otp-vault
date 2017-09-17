package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/99designs/keyring"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	KeyringName = "otp-vault"
)

var (
	keyringImpl       keyring.Keyring
	backendsAvailable = keyring.SupportedBackends()
)

var GlobalFlags struct {
	Debug        bool
	Backend      string
	PromptDriver string
}

func ConfigureGlobals(app *kingpin.Application) {
	app.Flag("debug", "Show debugging output").
		BoolVar(&GlobalFlags.Debug)

	app.Flag("backend", fmt.Sprintf("Secret backend to use %v", backendsAvailable)).
		Default(keyring.DefaultBackend).
		OverrideDefaultFromEnvar("OTP_VAULT_BACKEND").
		EnumVar(&GlobalFlags.Backend, backendsAvailable...)

	app.PreAction(func(c *kingpin.ParseContext) (err error) {
		if !GlobalFlags.Debug {
			log.SetOutput(ioutil.Discard)
		}
		if keyringImpl == nil {
			keyringImpl, err = keyring.Open(KeyringName, GlobalFlags.Backend)
		}
		return err
	})

}

func TerminalPrompt(prompt string) (string, error) {
	fmt.Fprint(os.Stderr, prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
	        return "", err
	}
	return strings.TrimSpace(text), nil
}

