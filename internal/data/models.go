package data

import (
	"database/sql"
)

type DBModels struct {
	Users   UserModel
	Doctors DoctorModel
}

func NewModels(db *sql.DB) DBModels {
	return DBModels{
		Users:   UserModel{},
		Doctors: DoctorModel{},
	}
}
