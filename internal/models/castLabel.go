package models

type CastLabel struct {
	ID    int32  `db:"movie_id"`
	Label string `db:"label"`
}
