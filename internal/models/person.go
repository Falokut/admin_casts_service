package models

type Person struct {
	ID           int32 `json:"id" db:"person_id"`
	ProfessionID int32 `json:"profession_id" db:"profession_id"`
}
