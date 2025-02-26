package bootstrap

import (
	"bytes"
	"os"
	"strings"

	cfg "github.com/cndrsdrmn/crugo/config"
	"github.com/spf13/viper"
)

type resource struct {
	cfg.AppConfig `mapstructure:"app"`
	cfg.DBConfig  `mapstructure:"db"`
}

func (app *App) bootConfiguration() error {
	data, err := os.ReadFile(app.ConfigPath)
	if err != nil {
		return writeError(err, "read_file", app.ConfigPath)
	}

	rvn := viper.New()
	rvn.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	rvn.AutomaticEnv()
	rvn.SetConfigType("yaml")

	if err := rvn.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return writeError(err, "read_config", app.ConfigPath)
	}

	var src resource

	if err := rvn.Unmarshal(&src); err != nil {
		return writeError(err, "unmarshal_config")
	}

	cfg.App = &src.AppConfig
	cfg.DB = &src.DBConfig

	return nil
}
