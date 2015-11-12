package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/codegangsta/cli"

	"github.com/benwebber/mipper/amo"
)

func Info(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal(errors.New("you must provide an addon name or ID"))
	}

	fmt.Sprintf("+%v", ctx)

	amoClient := amo.NewAMOClient()
	addon, err := amoClient.AddonByIdOrName(ctx.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %v\n", addon.Name)
	fmt.Printf("Version: %v\n", addon.Version)
	fmt.Printf("Homepage: %v\n", addon.Homepage)
	fmt.Printf("Summary: %v\n", addon.Summary)
}
