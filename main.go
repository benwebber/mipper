package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/benwebber/mipper/commands"
)

const (
	Version = "0.1.0"

	descGeneral = "Build multiple item packages for Firefox."
	descAdd     = "Add an addon to a package"
	descBuild   = "Build a multiple item package"
	descInfo    = "Show information about an addon"
	descRemove  = "Remove an addon from a package"
	descSearch  = "Search for addons"
	descVersion = "Print the version and exit"
)

func init() {
	// Override the default top-level and command help templates.
	cli.AppHelpTemplate = `Usage: {{.Name}} <command>{{if .Flags}} [options]{{end}} <args>...

  {{.Usage}}
{{if .Flags}}
Options:
   {{range .Flags}}{{.}}
   {{end}}
{{end}}Commands:
   {{range .Commands}}{{.Name}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}{{end}}
`
	cli.CommandHelpTemplate = `Usage: mipper {{.Name}}{{if .Flags}} [options]{{end}} <args>...
{{if .Description}}
  {{.Description}}
{{end}}{{if .Flags}}
Options:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`
}

func main() {
	app := cli.NewApp()
	app.Name = "mipper"
	app.Usage = descGeneral
	app.Version = Version
	app.HideVersion = true
	app.Commands = []cli.Command{
		{
			Name:        "add",
			ShortName:   "a",
			Usage:       descAdd,
			Description: descAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "manifest, f",
					Usage: "Path to package manifest",
				},
			},
			Action: commands.Add,
		},
		{
			Name:      "build",
			ShortName: "b",
			Usage:     descBuild,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "Path to package",
				},
			},
			Action: commands.Build,
		},
		{
			Name:      "info",
			ShortName: "i",
			Usage:     descInfo,
			Action:    commands.Info,
		},
		{
			Name:      "remove",
			ShortName: "r",
			Usage:     descRemove,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "manifest, f",
					Usage: "Path to package manifest",
				},
			},
			Action: commands.Remove,
		},
		{
			Name:      "search",
			ShortName: "s",
			Usage:     descSearch,
			Action:    commands.Search,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "headers, H",
					Usage: "Show table headers",
				},
			},
		},
		{
			Name:      "version",
			ShortName: "v",
			Usage:     descVersion,
			Action:    commands.Version,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
