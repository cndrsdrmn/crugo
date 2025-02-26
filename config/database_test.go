package config

import (
	"errors"
	"fmt"
	"testing"

	driver "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

func TestConnector(t *testing.T) {
	t.Run("should return sqlite connector", func(t *testing.T) {
		cfg := &DBConfig{
			Default:  "sqlite",
			Database: ":memory:",
		}

		conn, err := cfg.Connector()

		expected := sqlite.Open(cfg.Database)

		assert.NoError(t, err)
		assert.Equal(t, expected, conn)
		assert.Equal(t, "sqlite", conn.Name())
	})

	t.Run("should return mysql connector", func(t *testing.T) {
		cfg := &DBConfig{
			Default:  "mysql",
			Database: "crugo",
			Username: "crugo",
			Password: "crugo",
			Host:     "localhost",
			Port:     3306,
		}

		conn, err := cfg.Connector()

		expected := mysql.New(mysql.Config{
			DSNConfig: &driver.Config{
				User:      cfg.Username,
				Passwd:    cfg.Password,
				DBName:    cfg.Database,
				Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
				Net:       "tcp",
				ParseTime: true,
				Params: map[string]string{
					"charset": "utf8mb4",
				},
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, expected, conn)
		assert.Equal(t, "mysql", conn.Name())
	})

	t.Run("should return postgres connector", func(t *testing.T) {
		cfg := &DBConfig{
			Default:  "pgsql",
			Database: "crugo",
			Username: "crugo",
			Password: "crugo",
			Host:     "localhost",
			Port:     5432,
		}

		conn, err := cfg.Connector()

		expected := postgres.New(postgres.Config{
			DSN: fmt.Sprintf(
				"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
				cfg.Username, cfg.Password, cfg.Database, cfg.Host, cfg.Port,
			),
		})

		assert.NoError(t, err)
		assert.Equal(t, expected, conn)
		assert.Equal(t, "postgres", conn.Name())
	})

	t.Run("should return unsupported connector", func(t *testing.T) {
		cfg := &DBConfig{
			Default: "mongodb",
		}

		_, err := cfg.Connector()

		assert.EqualError(t, err, errors.ErrUnsupported.Error())
	})
}
