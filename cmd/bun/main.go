package main

import (
	"context"
	"fmt"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"os"
	"rps/cmd/bun/migrations"
	"rps/config"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	db := config.InitDatabase(cfg)

	if err != nil {
		panic(err)
	}

	app := &cli.App{
		Name:  "Bun migrations",
		Usage: "rock paper scissors",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Create migration tables",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)
					return migrator.Init(ctx)
				},
			},
			{
				Name:  "migrate",
				Usage: "Migrate database",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)

					group, err := migrator.Migrate(ctx)
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Println("There are no new migrations to run")
						return nil
					}

					fmt.Printf("Migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "Rollback the last migration group",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)

					group, err := migrator.Rollback(ctx)
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Println("There are no groups to roll back")
						return nil
					}

					fmt.Printf("Rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "Lock migrations",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)
					return migrator.Lock(ctx)
				},
			},
			{
				Name:  "unlock",
				Usage: "Unlock migrations",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)
					return migrator.Unlock(ctx)
				},
			},
			{
				Name:  "create_go",
				Usage: "Create Go migration",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)

					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(ctx, name)
					if err != nil {
						return err
					}
					fmt.Printf("Created migration %s (%s)\n", mf.Name, mf.Path)

					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "Create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)

					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(ctx, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("Created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "Print migrations status",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)

					ms, err := migrator.MigrationsWithStatus(ctx)
					if err != nil {
						return err
					}
					fmt.Printf("Migrations: %s\n", ms)
					fmt.Printf("Unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("Last migration group: %s\n", ms.LastGroup())

					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "Mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					ctx := context.Background()

					migrator := migrate.NewMigrator(db, migrations.Migrations)

					group, err := migrator.Migrate(ctx, migrate.WithNopMigration())
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Println("There are no new migrations to mark as applied")
						return nil
					}

					fmt.Printf("Marked as applied %s\n", group)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
