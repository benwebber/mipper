package commands

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/benwebber/mipper/pkg"
)

func Build(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal(errors.New("you must specify a package manifest"))
	}
	manifest := ctx.Args()[0]

	filename := ctx.String("output")
	// The user did not specify an output file. Default to <manifest>.xpi.
	if filename == "" {
		extension := filepath.Ext(manifest)
		filename = fmt.Sprintf("%v.xpi", manifest[0:len(manifest)-len(extension)])
	}

	p, err := pkg.NewFromFile(manifest)
	if err != nil {
		log.Fatal(err)
	}
	p.Build(filename)
}
