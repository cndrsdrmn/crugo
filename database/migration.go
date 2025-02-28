package database

import (
	"github.com/cndrsdrmn/crugo/facades"
	"github.com/cndrsdrmn/crugo/users"
)

func MigrateAndSeed() {
	migrate()
	seeder()
}

func migrate() {
	facades.DB.AutoMigrate(
		&users.User{},
	)
}

func seeder() {
	users := []*users.User{
		{Name: "Foo", Email: "foo@crugo.com", Password: "secret"},
		{Name: "Bar", Email: "bar@crugo.com", Password: "secret"},
	}

	facades.DB.Create(&users)
}
