package db

import (
	"github.com/raymondgitonga/company-service/internal/model"
)

type Person struct{}

type IPerson interface {
	GetPerson(email string) (model.Person, error)
}

func (p Person) GetPerson(email string) (model.Person, error) {
	var user model.Person
	db := createDBConnection()

	defer db.Close()

	query := `SELECT * FROM person WHERE email = $1;`

	row := db.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Email)

	return user, err
}
