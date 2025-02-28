package bootstrap_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/cndrsdrmn/crugo/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestBootstrap(t *testing.T) {
	cfg := []byte(`
app:
  name: crugo
  port: 8080
  debug: true
db:
  connection: sqlite
  database: ":memory:"
`)
	dbCfg := []byte(`
app:
  name: crugo
  port: 8080
  debug: true
db:
  connection: mongodb
`)
	os.WriteFile("/tmp/config.yaml", cfg, 0644)
	os.WriteFile("/tmp/database-config.yaml", dbCfg, 0644)
	defer os.Remove("/tmp/config.yaml")
	defer os.Remove("/tmp/database-config.yaml")

	t.Run("successfully bootstraping", func(t *testing.T) {
		app := &App{
			ConfigPath: "/tmp/config.yaml",
		}

		assert.False(t, app.IsBootstrapped)

		app.Bootstrap()

		assert.True(t, app.IsBootstrapped)
	})

	t.Run("fail bootstraping configuration", func(t *testing.T) {
		app := &App{
			ConfigPath: "/tmp/missing-config.yml",
		}

		assert.False(t, app.IsBootstrapped)

		err := app.Bootstrap()

		assert.EqualError(t, err, fmt.Sprintf("can't open %s: no such file or directory", app.ConfigPath))
		assert.False(t, app.IsBootstrapped)
	})

	t.Run("fail bootstrapping database", func(t *testing.T) {
		app := &App{
			ConfigPath: "/tmp/database-config.yaml",
		}

		assert.False(t, app.IsBootstrapped)

		err := app.Bootstrap()

		assert.EqualError(t, err, "unsupported database mongodb")
		assert.False(t, app.IsBootstrapped)
	})
}
