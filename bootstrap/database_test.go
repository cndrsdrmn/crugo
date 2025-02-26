package bootstrap

import (
	"fmt"
	"testing"

	"github.com/cndrsdrmn/crugo/config"
	"github.com/cndrsdrmn/crugo/db"
	"github.com/stretchr/testify/assert"
)

func TestBootDatabase(t *testing.T) {
	app := &App{}

	config.DB = &config.DBConfig{
		Username: "crugo",
		Password: "secret",
		Host:     "localhost",
	}

	config.App = &config.AppConfig{
		Debug: true,
	}

	t.Run("can connect with sqlite", func(t *testing.T) {
		config.DB.Default = config.SQLite
		config.DB.Database = ":memory:"

		err := app.bootDatabase()

		assert.NoError(t, err)
		assert.Equal(t, "sqlite", db.Instance.Name())
	})

	t.Run("can connect with mysql", func(t *testing.T) {
		config.DB.Default = config.MySQL
		config.DB.Database = "crugo"
		config.DB.Port = 3306

		err := app.bootDatabase()

		assert.NoError(t, err)
		assert.Equal(t, "mysql", db.Instance.Name())
	})

	t.Run("can't connect with missing database mysql", func(t *testing.T) {
		config.DB.Default = config.MySQL
		config.DB.Database = "missing"
		config.DB.Port = 3306

		err := app.bootDatabase()

		assert.EqualError(t, err, fmt.Sprintf("failed connect to database: %v", config.MySQL))
	})

	t.Run("can connect with postgres", func(t *testing.T) {
		config.DB.Default = config.Postgres
		config.DB.Database = "crugo"
		config.DB.Port = 5432

		err := app.bootDatabase()

		assert.NoError(t, err)
		assert.Equal(t, "postgres", db.Instance.Name())
	})

	t.Run("can't connect with missing database postgres", func(t *testing.T) {
		config.DB.Default = config.Postgres
		config.DB.Database = "missing"
		config.DB.Port = 5432

		err := app.bootDatabase()

		assert.EqualError(t, err, fmt.Sprintf("failed connect to database: %v", config.Postgres))
	})

	t.Run("can't connect with unsupported database", func(t *testing.T) {
		config.DB.Default = "mongodb"

		err := app.bootDatabase()

		assert.EqualError(t, err, "unsupported database mongodb")
	})
}
