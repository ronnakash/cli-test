package main

import (
	"os"

	"github.com/docker/docker/client"
	"github.com/rivo/tview"
)

type Config struct {
	ProjectDir string `mapstructure:"project_dir"`
}

var app *tview.Application
var config Config = Config{
	ProjectDir: "~/repositories/java",
}
var dockerClient *client.Client
var checksums map[string]string

func run() int {
	ctx, logger, cancel := NewLogContext()

	logger.Info("starting cli app")
	defer cancel()

	app, err := NewCliApplication()
	if err != nil {
		logger.
			With("error", err).
			Info("failed to init app")
		return 1
	}

	if err := app.Run(); err != nil {
		logger.
			With("error", err).
			Info("failed to run app")
		return 1
	}

	return 0
}

func main() {
	exitCode := run()
	os.Exit(exitCode)
}
