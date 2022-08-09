package db

import (
	"github.com/raymondgitonga/company-service/internal/config"
	"github.com/raymondgitonga/company-service/internal/model"
)

type Person struct {
	email string
	role  string
}

func NewPerson(email string) iPerson {
	return &Person{email: email}
}

type iPerson interface {
	GetPerson() (model.Person, error)
}

func (p Person) GetPerson() (model.Person, error) {
	var person model.Person
	db := config.CreateDBConnection()

	defer db.Close()

	query := `SELECT * FROM person WHERE email = $1;`

	row := db.QueryRow(query, p.email)

	err := row.Scan(&person.Email, &person.Email, &person.Role)

	return person, err
}
