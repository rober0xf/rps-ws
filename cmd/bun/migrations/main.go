package migrations

import (
	"embed"
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
