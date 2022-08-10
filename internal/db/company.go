package db

import (
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

func NewCompany(name string, code string, country string, website string, phone string) ICompany {
	return &Company{
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
}

func (c *Company) GetCompanies() ([]Company, error) {
	var companies []Company
	db := config.CreateDBConnection()

	defer db.Close()

	query := "SELECT * FROM company"

	queryString := c.buildQueryString(query)

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

	queryString := c.buildQueryString(query)

	row := db.QueryRow(queryString)

	err := row.Scan(&company.ID, &company.Name, &company.Code, &company.Country, &company.Website, &company.Phone)

	if err != nil {
		return Company{}, err
	}

	return company, nil

}

func (c *Company) buildQueryString(query string) string {
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
