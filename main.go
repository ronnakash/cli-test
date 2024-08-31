package main

import (
	"os"
)

func run() int {

	app, err := NewCliApplication()
	if err != nil {
		return 1
	}

	if err := app.Run(); err != nil {
		return 1
	}

	return 0
}

func main() {
	exitCode := run()
	os.Exit(exitCode)
}
