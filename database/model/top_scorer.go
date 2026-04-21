package model

type TopScorer struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Club     string `db:"club"`
	Position string `db:"position"`
	Goals    int    `db:"goals"`
}
