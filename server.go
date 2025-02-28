package main

import (
	"github.com/cndrsdrmn/crugo/bootstrap"
)

type Foo struct {
	Bar     string `json:"bar" validate:"required,min=4"`
	Email   string `json:"email" validate:"required,email"`
	FooBar  string `json:"foo_bar" validate:"required"`
	Numeric int
}

func main() {
	app := &bootstrap.App{
		ConfigPath: "config.yaml",
	}

	app.Serve()
}
