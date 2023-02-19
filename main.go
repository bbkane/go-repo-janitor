package main

import (
	"os"

	"go.bbkane.com/warg"
	"go.bbkane.com/warg/section"
)

var version string

func buildApp() warg.App {
	app := warg.New(
		"go-repo-janitor",
		section.New(
			"Do common Go repo automation tasks",
			section.Command(
				"hello",
				"Say hello",
				hello,
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
