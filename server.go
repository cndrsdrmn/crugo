package main

import (
	"github.com/cndrsdrmn/crugo/bootstrap"
)

func main() {
	app := &bootstrap.App{
		ConfigPath: "config.yaml",
	}

	app.Serve()
}
