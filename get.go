package main

import (
	"encoding/json"
	"fmt"
	"time"
	"strconv"

	"github.com/99designs/keyring"
	"github.com/hgfischer/go-otp"
	"gopkg.in/alecthomas/kingpin.v2"
)

type GetCommandInput struct {
	Profile string
}

func ConfigureGetCommand(app *kingpin.Application) {
	input := GetCommandInput{}

	cmd := app.Command("get", "Gets an OTP key for a profile.")

	cmd.Arg("Profile", "Name of the profile to get a key for.").
		Required().
		StringVar(&input.Profile)

	cmd.Action(func(c *kingpin.ParseContext) error {
		GetCommand(app, input, keyringImpl)
		return nil
	})
}

func GetCommand(app *kingpin.Application, input GetCommandInput, keyringImpl keyring.Keyring) {
	var code string
	var context string
	var err error
	var obj OTPParameters

	item, err := keyringImpl.Get(input.Profile)

	if err != nil {
		app.Fatalf(err.Error())
		return
	}

	err = json.Unmarshal(item.Data, &obj)

	if err != nil {
		app.Fatalf(err.Error())
		return
	}

	if obj.Type == "hotp" {
		hotp := &otp.HOTP{
			Counter: uint64(obj.Counter),
			IsBase32Secret: true,
			Length: uint8(obj.Length),
			Secret: obj.Secret,
		}
		code = hotp.Get()
		context =  strconv.Itoa(obj.Counter)

		obj.Counter = obj.Counter + 1
		bytes, err := json.Marshal(obj)

		if err != nil {
			app.Fatalf(err.Error())
			return
		}

		err = keyringImpl.Set(
			keyring.Item{
				Key: input.Profile,
				Data: bytes,
			})

	} else {
		totp := &otp.TOTP{
			IsBase32Secret: true,
			Length: uint8(obj.Length),
			Period: uint8(obj.Period),
			Secret: obj.Secret,
		}
		code = totp.Get()
		context = time.Now().String()
	}

	fmt.Printf("%s (%s)\n", code, context)
}
