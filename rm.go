package main

import (
	"fmt"

	"github.com/99designs/keyring"
	"gopkg.in/alecthomas/kingpin.v2"
)

type RemoveCommandInput struct {
	Profile      string
	Keyring      keyring.Keyring
}

func ConfigureRemoveCommand(app *kingpin.Application) {
	input := RemoveCommandInput{}

	cmd := app.Command("remove", "Removes credentials")
	cmd.Alias("rm")

	cmd.Arg("profile", "Name of the profile").
		Required().
		StringVar(&input.Profile)

	cmd.Action(func(c *kingpin.ParseContext) error {
		input.Keyring = keyringImpl
		RemoveCommand(app, input)
		return nil
	})
}

func RemoveCommand(app *kingpin.Application, input RemoveCommandInput) {
	r, err := TerminalPrompt(fmt.Sprintf("Delete credentials for profile %q? (Y|n)", input.Profile))
	if err != nil {
		app.Fatalf(err.Error())
		return
	} else if r == "N" || r == "n" {
		return
	}

	if err := input.Keyring.Remove(input.Profile); err != nil {
		app.Fatalf(err.Error())
		return
	}
	fmt.Printf("Deleted credentials.\n")
}
