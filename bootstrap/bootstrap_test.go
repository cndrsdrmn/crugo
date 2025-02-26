package bootstrap

import (
	"os"
	"testing"

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

		assert.False(t, app.isBootstrapped)

		app.Bootstrap()

		assert.True(t, app.isBootstrapped)
	})

	t.Run("fail bootstraping configuration", func(t *testing.T) {
		app := &App{
			ConfigPath: "/tmp/missing-config.yml",
		}

		assert.False(t, app.isBootstrapped)

		app.Bootstrap()

		assert.False(t, app.isBootstrapped)
	})

	t.Run("fail bootstrapping database", func(t *testing.T) {
		app := &App{
			ConfigPath: "/tmp/database-config.yaml",
		}

		assert.False(t, app.isBootstrapped)

		app.Bootstrap()

		assert.False(t, app.isBootstrapped)
	})
}
