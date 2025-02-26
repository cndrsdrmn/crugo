package bootstrap

import (
	"fmt"
	"log"
)

type App struct {
	ConfigPath     string
	isBootstrapped bool
}

func (app *App) Bootstrap() error {
	if err := app.bootConfiguration(); err != nil {
		return err
	}

	if err := app.bootDatabase(); err != nil {
		return err
	}

	app.isBootstrapped = true

	return nil
}

func writeError(err error, on string, args ...any) error {
	log.Default().Println(err)

	switch on {
	case "read_file":
		err = fmt.Errorf("can't open %v: no such file or directory", args...)
	case "read_config":
		err = fmt.Errorf("can't parsing config: %v", args...)
	case "unmarshal_config":
		err = fmt.Errorf("can't unmarshal config")
	case "db_connection":
		err = fmt.Errorf("failed connect to database: %v", args...)
	case "db_unsupported":
		err = fmt.Errorf("unsupported database %v", args...)
	}

	return err
}
