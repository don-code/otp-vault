package main

import (
	"encoding/json"
	"fmt"

	"github.com/99designs/keyring"
	"gopkg.in/alecthomas/kingpin.v2"
)

type OTPParameters struct {
	Counter int
	Length int
	Period int
	Profile string
	Secret string
	Type string
}

func ConfigureAddCommand(app *kingpin.Application) {
	input := OTPParameters{}

	cmd := app.Command("add", "Adds a profile")

	cmd.Arg("Profile", "Name of the profile").
		Required().
		StringVar(&input.Profile)

	cmd.Arg("Secret", "Secret key, as base32").
		Required().
		StringVar(&input.Secret)

	cmd.Arg("Type", "The type of OTP password, either 'totp' or 'hotp'. Default 'totp'.").
		Default("totp").
		StringVar(&input.Type)

	cmd.Arg("Counter", "The counter used as a moving factor (HOTP only). Default 0.").
		Default("0").
		IntVar(&input.Counter)

	cmd.Arg("Period", "The step size to slice time, in seconds (TOTP only). Default 30.").
		Default("30").
		IntVar(&input.Period)

	cmd.Arg("Length", "Token length. Default 6.").
		Default("6").
		IntVar(&input.Length)

	cmd.Action(func(c *kingpin.ParseContext) error {
		AddCommand(app, input, keyringImpl)
		return nil
	})
}

// TODO: input sanitization
func AddCommand(app *kingpin.Application, input OTPParameters, keyringImpl keyring.Keyring) {
	var err error

	if !(input.Type == "hotp" || input.Type == "totp") {
		app.Fatalf("Type field must be either 'hotp' or 'totp'.")
		return
	}

	bytes, err := json.Marshal(input)

	if err != nil {
		app.Fatalf(err.Error())
		return
	}

	err = keyringImpl.Set(
		keyring.Item{
			Key: input.Profile,
			Data: bytes,
		})

	if err != nil {
		app.Fatalf(err.Error())
		return
	}

	fmt.Printf("Set up %q in vault\n", input.Profile)

}
