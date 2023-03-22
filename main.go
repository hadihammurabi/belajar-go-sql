package main

import (
	"belajar-go-sqlx/config"
	"belajar-go-sqlx/create"
	"belajar-go-sqlx/insert"
	sel "belajar-go-sqlx/select"

	"github.com/jmoiron/sqlx"
)

func main() {
	config.DB = sqlx.MustOpen("sqlite3", ":memory:")
	create.Migration()
	insert.Insert()
	sel.Select()
	sel.SelectScan()
	sel.SelectScanV2()
}
