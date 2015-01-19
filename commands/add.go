package commands

import "github.com/codegangsta/cli"

func Add(ctx *cli.Context) {
	modifyPackage(ctx, "add")
}
