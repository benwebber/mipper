package commands

import "github.com/codegangsta/cli"

func Remove(ctx *cli.Context) {
	modifyPackage(ctx, "remove")
}
