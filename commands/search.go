package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, addon := range addons {
		fmt.Fprintf(w, "%v\t%s\n", addon.ID, addon.Name)
	}
	w.Flush()
}
