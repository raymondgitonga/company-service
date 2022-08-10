package db

import (
	"errors"
	"fmt"
	"github.com/raymondgitonga/company-service/internal/config"
	"strings"
)

type Company struct {
	ID      int    `json:"companyId"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
	//Email   string `json:"email"`
}

func NewCompany(id int, name string, code string, country string, website string, phone string) ICompany {
	return &Company{
		ID:      id,
		Name:    name,
		Code:    code,
		Country: country,
		Website: website,
		Phone:   phone,
		//Email:   email,
	}
}

type ICompany interface {
	GetCompanies() ([]Company, error)
	GetCompany() (Company, error)
	CreateCompany() error
	DeleteCompany() error
	UpdateCompany(id string) error
}

func (c *Company) GetCompanies() ([]Company, error) {
	var companies []Company
	db := config.CreateDBConnection()

	defer db.Close()

	query := "SELECT * FROM company"

	queryString := c.buildSearchQuery(query)

	rows, err := db.Query(queryString)

	if err != nil {
		return []Company{}, err
	}

	for rows.Next() {
		var company Company

		err := rows.Scan(&company.ID, &company.Name, &company.Code, &company.Country, &company.Website, &company.Phone)

		if err != nil {
			return []Company{}, err
		}
		companies = append(companies, company)
	}

	return companies, err
}

func (c *Company) GetCompany() (Company, error) {
	var company Company
	db := config.CreateDBConnection()

	defer db.Close()

	query := "SELECT * FROM company"

	queryString := c.buildSearchQuery(query)

	row := db.QueryRow(queryString)

	err := row.Scan(&company.ID, &company.Name, &company.Code, &company.Country, &company.Website, &company.Phone)

	if err != nil {
		return Company{}, err
	}

	return company, nil

}

func (c *Company) CreateCompany() error {
	db := config.CreateDBConnection()
	defer db.Close()

	query := `INSERT INTO company (name, code, country, website, phone) VALUES ($1, $2, $3, $4, $5)`

	err := db.QueryRow(query, c.Name, c.Code, c.Country, c.Website, c.Phone)

	if err != nil {
		return err.Err()
	}

	return nil
}

func (c *Company) DeleteCompany() error {

	db := config.CreateDBConnection()
	defer db.Close()

	query := `DELETE FROM company WHERE companyId = $1;`

	res, err := db.Exec(query, c.ID)
	if err != nil {
		return err
	}
	var row int64 = 0

	row, _ = res.RowsAffected()

	if row < 1 {
		return errors.New("failed to delete company")
	}

	return nil
}

func (c *Company) UpdateCompany(id string) error {
	db := config.CreateDBConnection()

	defer db.Close()

	// UPDATE company set code = 125 where companyid = 4;

	query := "UPDATE company SET"

	queryString, err := c.buildUpdateQuery(query, id)

	res, err := db.Exec(queryString)

	if err != nil {
		return err
	}

	rows, rowErr := res.RowsAffected()

	if rows < 1 {
		return rowErr
	}

	return nil
}

func (c Company) buildUpdateQuery(query string, id string) (string, error) {
	ext := make([]string, 0)
	if c.Name != "" {
		ext = append(ext, fmt.Sprintf("name = '%s'", c.Name))
	}

	if c.Phone != "" {
		ext = append(ext, fmt.Sprintf("phone = '%s'", c.Phone))
	}

	if c.Code != "" {
		ext = append(ext, fmt.Sprintf("code = '%s'", c.Code))
	}

	if c.Website != "" {
		ext = append(ext, fmt.Sprintf("website = '%s'", c.Website))
	}

	if c.Country != "" {
		ext = append(ext, fmt.Sprintf("country = '%s'", c.Country))
	}

	if len(ext) > 0 {
		return fmt.Sprintf("%s  %s WHERE companyid = %s", query, strings.Join(ext, " , "), id), nil
	}

	return "", errors.New("no update parameters")
}

func (c *Company) buildSearchQuery(query string) string {
	ext := make([]string, 0)
	if c.Name != "" {
		ext = append(ext, fmt.Sprintf("name = '%s'", c.Name))
	}

	if c.Phone != "" {
		ext = append(ext, fmt.Sprintf("phone = '%s'", c.Phone))
	}

	if c.Code != "" {
		ext = append(ext, fmt.Sprintf("code = '%s'", c.Code))
	}

	if c.Website != "" {
		ext = append(ext, fmt.Sprintf("website = '%s'", c.Website))
	}

	if c.Country != "" {
		ext = append(ext, fmt.Sprintf("country = '%s'", c.Country))
	}

	if len(ext) > 0 {
		return fmt.Sprintf("%s WHERE %s", query, strings.Join(ext, " and "))
	}

	return query
}
