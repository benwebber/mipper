package commands

import (
	"errors"
	"log"

	"github.com/codegangsta/cli"

	"github.com/benwebber/mipper/amo"
	"github.com/benwebber/mipper/pkg"
)

func modifyPackage(ctx *cli.Context, action string) {

	actions := map[string]func(p *pkg.Package, addon *amo.Addon){
		"add": func(p *pkg.Package, addon *amo.Addon) {
			p.Add(*addon)
		},
		"remove": func(p *pkg.Package, addon *amo.Addon) {
			p.Remove(*addon)
		},
	}

	manifest := ctx.String("manifest")

	if len(ctx.Args()) != 1 {
		log.Fatal(errors.New("you must provide an addon name or ID"))
	}

	if manifest == "" {
		log.Fatal(errors.New("you must provide a package manifest"))
	}

	p, err := pkg.NewFromFile(manifest)
	if err != nil {
		log.Fatal(err)
	}

	amoClient := amo.NewAMOClient()
	addon, err := amoClient.AddonByIdOrName(ctx.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	actions[action](p, addon)
	p.WriteFile(manifest)
}
