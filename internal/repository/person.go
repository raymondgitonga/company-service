package repository

import (
	"github.com/raymondgitonga/company-service/internal/config"
	"github.com/raymondgitonga/company-service/internal/model"
)

type Person struct {
	Email string
}

func NewPerson(email string) PersonRepository {
	return &Person{Email: email}
}

type PersonRepository interface {
	GetPerson() (model.Person, error)
}

func (p Person) GetPerson() (model.Person, error) {
	var person model.Person
	db := config.CreateDBConnection()

	defer db.Close()

	query := `SELECT * FROM person WHERE Email = $1;`

	row := db.QueryRow(query, p.Email)

	err := row.Scan(&person.Email, &person.Email, &person.Role)

	return person, err
}
