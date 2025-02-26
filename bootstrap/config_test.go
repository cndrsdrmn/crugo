package bootstrap

import (
	"fmt"
	"os"
	"testing"

	"github.com/cndrsdrmn/crugo/config"
	"github.com/stretchr/testify/assert"
)

func TestBootConfiguration(t *testing.T) {
	invalid := []byte(`app: port: 8080`)
	unmarshal := []byte(`app: "should_be_a_map"`)
	valid := []byte(`
app:
  name: crugo
  port: 8080
  debug: true
db:
  connection: sqlite
  database: ":memory:"
`)

	os.WriteFile("/tmp/config.yaml", valid, 0644)
	os.WriteFile("/tmp/unmarshal-config.yaml", unmarshal, 0644)
	os.WriteFile("/tmp/invalid-config.yaml", invalid, 0644)

	defer os.Remove("/tmp/config.yaml")
	defer os.Remove("/tmp/invalid-config.yaml")

	app := &App{
		ConfigPath: "/tmp/config.yaml",
	}

	t.Run("read configuration from file", func(t *testing.T) {
		app.bootConfiguration()

		assert.Equal(t, "crugo", config.App.Name)
		assert.Equal(t, 8080, config.App.Port)
		assert.Equal(t, true, config.App.Debug)
		assert.Equal(t, config.DBConnection("sqlite"), config.DB.Default)
		assert.Equal(t, ":memory:", config.DB.Database)
	})

	t.Run("override configuration from environment", func(t *testing.T) {
		defer os.Clearenv()
		os.Setenv("APP_NAME", "crugo-env")
		os.Setenv("APP_PORT", "9090")
		os.Setenv("APP_DEBUG", "false")
		os.Setenv("DB_CONNECTION", "mysql")
		os.Setenv("DB_DATABASE", "crugo-env.db")

		app.bootConfiguration()

		assert.Equal(t, "crugo-env", config.App.Name)
		assert.Equal(t, 9090, config.App.Port)
		assert.Equal(t, false, config.App.Debug)
		assert.Equal(t, config.DBConnection("mysql"), config.DB.Default)
		assert.Equal(t, "crugo-env.db", config.DB.Database)
	})

	t.Run("fail to read configuration from file", func(t *testing.T) {
		app.ConfigPath = "/tmp/config-not-exist.yaml"

		err := app.bootConfiguration()

		assert.EqualError(t, err, fmt.Sprintf("can't open %v: no such file or directory", app.ConfigPath))
	})

	t.Run("fail to read configuration from invalid file", func(t *testing.T) {
		app.ConfigPath = "/tmp/invalid-config.yaml"

		err := app.bootConfiguration()

		assert.EqualError(t, err, fmt.Sprintf("can't parsing config: %v", app.ConfigPath))
	})

	t.Run("fail to unmarshal config", func(t *testing.T) {
		app.ConfigPath = "/tmp/unmarshal-config.yaml"

		err := app.bootConfiguration()

		assert.EqualError(t, err, "can't unmarshal config")
	})
}
