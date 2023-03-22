package sel

import (
	"belajar-go-sqlx/config"
	"fmt"
	"time"

	"github.com/blockloop/scan/v2"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       uint
	Email    string
	Password string
}

func Select() {
	rows, err := config.DB.Queryx(`SELECT * FROM users`)
	config.Must(err)

	users := make([]User, 0)

	start := time.Now()
	for rows.Next() {
		var user User
		if err := rows.StructScan(&user); err == nil {
			users = append(users, user)
		}
	}
	end := time.Since(start)
	fmt.Println("StructScan:", end)
}

func SelectScan() {
	rows, err := config.DB.Query(`SELECT * FROM users`)
	config.Must(err)

	users := make([]User, 0)

	start := time.Now()
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)
		if err == nil {
			users = append(users, user)
		}
	}
	end := time.Since(start)
	fmt.Println("Scan:", end)
}

func SelectScanV2() {
	rows, err := config.DB.Query(`SELECT * FROM users`)
	config.Must(err)

	users := make([]User, 0)

	start := time.Now()
	scan.Rows(&users, rows)
	end := time.Since(start)
	fmt.Println("Scan:", end)
}
