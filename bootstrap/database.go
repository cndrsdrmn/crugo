package bootstrap

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (app *App) bootDatabase() error {
	cfg := app.config

	conn, err := app.config.DBConfig.Connector()
	if err != nil {
		return writeError(err, "db_unsupported", cfg.DBConfig.Default)
	}

	gc := &gorm.Config{}
	if cfg.AppConfig.Debug {
		gc.Logger = logger.Default.LogMode(logger.Info)
	}

	app.db, err = gorm.Open(conn, gc)
	if err != nil {
		return writeError(err, "db_connection", conn.Name())
	}

	return nil
}
