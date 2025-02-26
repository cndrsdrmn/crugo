package bootstrap

import (
	"fmt"

	"github.com/cndrsdrmn/crugo/config"
	"github.com/cndrsdrmn/crugo/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (app *App) bootDatabase() error {
	conn, err := config.DB.Connector()
	if err != nil {
		return writeError(err, "db_unsupported", config.DB.Default)
	}

	gc := &gorm.Config{}
	if config.App.Debug {
		gc.Logger = logger.Default.LogMode(logger.Info)
	}

	db.Instance, err = gorm.Open(conn, gc)
	if err != nil {
		return writeError(err, "db_connection", conn.Name())
	}

	fmt.Println("database connected successfully")

	return nil
}
