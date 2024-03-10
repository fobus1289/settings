package main

import (
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	// go get -u github.com/fergusstrange/embedded-postgres
	// postgres := embeddedpostgres.NewDatabase()
	//
	// postgres.Start()
	// if err := postgres.Start(); err != nil {
	// 	slog.Error("error", err)
	// 	return
	// }

	// defer func() {
	// 	slog.Info("info", postgres.Stop())
	// }()

	pg, err := connect()
	if err != nil {
		slog.Error("error", err)
		return
	}

	pg.Get("", "", "")

	// defer func() {
	// 	pg.Close()
	// }()
	//
	// if err := goose.Up(pg.DB, "./migrations"); err != nil {
	// 	log.Fatal(err)
	// }
	// // goose_db_version
	// var i []string
	//
	// // if err := pg.Select(&i, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'"); err != nil {
	// // 	slog.Error("Error", err)
	// // 	return
	// // }
	//
	// if err := pg.Select(&i, "SELECT version_id FROM goose_db_version"); err != nil {
	// 	slog.Error("Error", err)
	// 	return
	// }
	//
	// slog.Info("info", slog.Any("idd", i))
}

func connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"pgx",
		"host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
	)

	return db, err
}
