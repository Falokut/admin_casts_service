package models

type Cast struct {
	Label          string `db:"label"`
	ProfessionName string `json:"profession_name" db:"profession_name"`
	ID             int32  `db:"movie_id"`
	PersonID       int32  `json:"person_id" db:"person_id"`
	ProfessionID   int32  `json:"profession_id" db:"profession_id"`
}
