package config

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
