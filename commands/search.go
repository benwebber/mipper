package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/codegangsta/cli"

	"github.com/benwebber/mipper/amo"
)

func Search(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal(errors.New("you must provide an addon name"))
	}
	addonName := ctx.Args()[0]

	amo := amo.NewAMOClient()
	addons, err := amo.Search(addonName)
	if err != nil {
		log.Fatal(err)
	}
	for _, addon := range addons {
		fmt.Println(addon.Name)
	}
}
