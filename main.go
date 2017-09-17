package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	run(os.Args[1:], os.Exit)
}

func run(args []string, exit func(int)) {
	app := kingpin.New(
		`otp-vault`,
		`A vault for securely retrieving one-time password (OTP) tokens.`,
	)

	app.Writer(os.Stdout)
	app.Version("0.1.0")
	app.Terminate(exit)

	ConfigureGlobals(app)
	ConfigureAddCommand(app)
	ConfigureGetCommand(app)
	ConfigureListCommand(app)
	ConfigureRemoveCommand(app)

	kingpin.MustParse(app.Parse(args))
}
