package db

import (
	"github.com/raymondgitonga/company-service/internal/config"
	"github.com/raymondgitonga/company-service/internal/model"
)

type Person struct {
	email string
}

func NewPerson(email string) iPerson {
	return &Person{email: email}
}

type iPerson interface {
	PersonExists() (bool, error)
}

func (p Person) PersonExists() (bool, error) {
	var person model.Person
	db := config.CreateDBConnection()

	defer db.Close()

	query := `SELECT * FROM person WHERE email = $1;`

	row := db.QueryRow(query, p.email)

	err := row.Scan(&person.Email, &person.Email)

	if len(person.Email) == 0 {
		return false, err
	}

	return true, err
}
