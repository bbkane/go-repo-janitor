package main

import (
	"os"

	"go.bbkane.com/warg"
	"go.bbkane.com/warg/command"
	"go.bbkane.com/warg/flag"
	"go.bbkane.com/warg/section"
	"go.bbkane.com/warg/value/scalar"
	"go.bbkane.com/warg/value/slice"
)

var version string

func buildApp() warg.App {
	app := warg.New(
		"go-repo-janitor",
		section.New(
			"Do common Go repo automation tasks",
			section.Command(
				"vimdiff",
				"Print vimdiff commands between repos",
				vimdiff,
				command.Flag(
					"--dst",
					"Destination repos to diff against",
					slice.Path(
						slice.Default(
							[]string{
								"~/Git-GH/fling",
								"~/Git-GH/go-repo-janitor",
								"~/Git-GH/gocolor",
								"~/Git-GH/grabbit",
								"~/Git-GH/logos",
								"~/Git-GH/starghaze",
								"~/Git-GH/tablegraph",
								"~/Git-GH/taggedmarks2",
								"~/Git-GH/warg",
							},
						),
					),
					flag.Alias("-d"),
					flag.ConfigPath("vimdiff.dst"),
					flag.Required(),
				),
				command.Flag(
					"--file",
					"Common file in both repos to diff against",
					scalar.Path(),
					flag.Alias("-f"),
					flag.ConfigPath("vimdiff.file"),
					flag.Required(),
				),
				command.Flag(
					"--src",
					"Source repo to diff against",
					scalar.Path(
						scalar.Default("~/Git-GH/example-go-cli"),
					),
					flag.Alias("-s"),
					flag.ConfigPath("vimdiff.src"),
					flag.Required(),
				),
			),
		),
		warg.AddColorFlag(),
		warg.AddVersionCommand(version),
	)
	return app
}

func main() {
	app := buildApp()
	app.MustRun(os.Args, os.LookupEnv)
}
