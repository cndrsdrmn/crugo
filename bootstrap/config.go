package bootstrap

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/viper"
)

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

	if err := rvn.Unmarshal(&app.config); err != nil {
		return writeError(err, "unmarshal_config")
	}

	return nil
}
