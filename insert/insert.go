package insert

import "belajar-go-sqlx/config"

func Insert() {
	tx := config.DB.MustBegin()
	tx.MustExec(`INSERT INTO users VALUES
		(
			ABS(RANDOM()),
			?,
			?
		), (
			ABS(RANDOM()),
			?,
			?
		)`,
		"alex@mail.com",
		"123123",
		"john@mail.com",
		"123123",
	)
	tx.Commit()
}
