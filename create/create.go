package create

import (
	"belajar-go-sqlx/config"

	_ "github.com/mattn/go-sqlite3"
)

func Migration() {
	config.DB.MustExec(`CREATE TABLE users(
		id ROWID NOT NULL,
		email VARCHAR(255) NOT NULL,
		password TEXT
	)`)
}
