package data

import (
	"database/sql"

	db "d.hospital/pkg"
)

type Doctor struct {
	name       string
	surname    string
	position   string
	age        string
	experience string
}

type DoctorModel struct {
	db *sql.DB
}

func NewDoctorRepository() *DoctorModel {
	return &DoctorModel{
		db: db.DB,
	}
}

func (r *DoctorModel) RegisterDr(name, surname, position, age, experience string) (string, error) {
	err := r.db.QueryRow("INSERT INTO doctors (name, surname, position, age, exp) VALUES ($1, $2, $3, $4, $5) RETURNING name", name, surname, position, age, experience).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (r *DoctorModel) ReadDr(name string) (*Doctor, error) {
	doctor := &Doctor{}
	err := r.db.QueryRow("GET * FROM doctors  WHERE name $1", name).Scan(&doctor.name, &doctor.surname, &doctor.position, &doctor.age, &doctor.experience)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func (r *DoctorModel) ListDr(name string) ([]*Doctor, error) {
	rows, err := r.db.Query("SELECT * FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []*Doctor
	for rows.Next() {
		doctor := &Doctor{}
		err = rows.Scan(&doctor.name, &doctor.surname, &doctor.position, &doctor.age, &doctor.experience)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func (r *DoctorModel) UpdateDr() {

}

func (r *DoctorModel) DeleteDr() {

}
