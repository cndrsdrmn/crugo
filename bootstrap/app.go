package bootstrap

import (
	"fmt"
	"log"

	cfg "github.com/cndrsdrmn/crugo/config"
	"github.com/cndrsdrmn/crugo/database"
	"github.com/cndrsdrmn/crugo/facades"
	"github.com/cndrsdrmn/crugo/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	ConfigPath     string
	IsBootstrapped bool
	Server         *gin.Engine
	config         appConfig
	db             *gorm.DB
}

type appConfig struct {
	cfg.AppConfig `mapstructure:"app"`
	cfg.DBConfig  `mapstructure:"db"`
}

func (app *App) Bootstrap() error {
	if err := app.bootConfiguration(); err != nil {
		return err
	}

	if err := app.bootDatabase(); err != nil {
		return err
	}

	app.afterBootstrapped()

	return nil
}

func (app *App) Serve() error {
	if err := app.ensureIsBootrapped(); err != nil {
		return err
	}
	// Run database migration and seeder
	database.MigrateAndSeed()
	// Register HTTP routes
	routes.RouteRegistar(app.Server)

	return app.Server.Run(fmt.Sprintf(":%v", cfg.App.Port))
}

func (app *App) afterBootstrapped() {
	app.IsBootstrapped = true

	// Populate configuration
	cfg.App = &app.config.AppConfig
	cfg.DB = &app.config.DBConfig

	// Populate facades
	facades.DB = app.db

	// Configure server
	mode := gin.ReleaseMode
	if cfg.App.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	app.Server = gin.Default()
}

func (app *App) ensureIsBootrapped() error {
	if app.IsBootstrapped {
		return nil
	}

	if err := app.Bootstrap(); err != nil {
		return err
	}

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
