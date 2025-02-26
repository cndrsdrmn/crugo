package config

import (
	"errors"
	"fmt"

	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *DBConfig

type DBConnection string

const (
	MySQL    DBConnection = "mysql"
	Postgres DBConnection = "postgres"
	SQLite   DBConnection = "sqlite"
)

type DBConfig struct {
	Default  DBConnection `mapstructure:"connection"`
	Database string       `yaml:"database"`
	Host     string       `yaml:"host"`
	Port     int          `yaml:"port"`
	Username string       `yaml:"username"`
	Password string       `yaml:"password"`
}

func (c *DBConfig) Connector() (gorm.Dialector, error) {
	var conn gorm.Dialector

	switch c.Default {
	case "mysql":
		conn = c.createMySQLConnector()
	case "pgsql", "postgres":
		conn = c.createPostgresConnector()
	case "sqlite":
		conn = c.createSQLiteConnector()
	default:
		return nil, errors.ErrUnsupported
	}

	return conn, nil
}

func (c *DBConfig) createMySQLConnector() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSNConfig: &driver.Config{
			User:      c.Username,
			Passwd:    c.Password,
			DBName:    c.Database,
			Addr:      fmt.Sprintf("%s:%d", c.Host, c.Port),
			Net:       "tcp",
			ParseTime: true,
			Params: map[string]string{
				"charset": "utf8mb4",
			},
		},
	})
}

func (c *DBConfig) createPostgresConnector() gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			c.Username, c.Password, c.Database, c.Host, c.Port,
		),
	})
}

func (c *DBConfig) createSQLiteConnector() gorm.Dialector {
	return sqlite.Open(c.Database)
}
