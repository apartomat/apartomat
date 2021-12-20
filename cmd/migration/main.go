package main

import (
	"flag"
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"os"
)

const usage = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.

Usage:
  migration <command> [args]
`

func main() {
	flag.Usage = func() {
		fmt.Print(usage)
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()

	pgopts, err := pg.ParseURL(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't parse POSTGRES_DSN %s", err)
		os.Exit(1)
	}

	db := pg.Connect(pgopts)

	migrations.SetTableName("public.migrations")

	prev, cur, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	if cur != prev {
		fmt.Printf("migrated from version %d to %d\n", prev, cur)
	} else {
		fmt.Printf("version is %d\n", cur)
	}
}
