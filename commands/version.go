package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func Version(ctx *cli.Context) {
	fmt.Printf("%v %v\n", ctx.App.Name, ctx.App.Version)
}
