package service

import (
	"github.com/raymondgitonga/company-service/internal/repository"
	"strconv"
)

type Company struct {
	ID      int    `json:"companyId"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

func NewCompany(id int, name string, code string, country string, website string, phone string) CompanyService {
	return &Company{
		ID:      id,
		Name:    name,
		Code:    code,
		Country: country,
		Website: website,
		Phone:   phone,
	}
}

type CompanyService interface {
	GetCompanies() ([]repository.Company, error)
	GetCompany() (repository.Company, error)
	CreateCompany() error
	DeleteCompany() error
	UpdateCompany(id string) error
}

func (c Company) GetCompanies() ([]repository.Company, error) {
	company := repository.NewCompany(c.ID, c.Name, c.Code, c.Country, c.Website, c.Phone)
	return company.GetCompanies()
}

func (c Company) GetCompany() (repository.Company, error) {
	company := repository.NewCompany(c.ID, c.Name, c.Code, c.Country, c.Website, c.Phone)
	return company.GetCompany()
}

func (c Company) CreateCompany() error {
	company := repository.NewCompany(c.ID, c.Name, c.Code, c.Country, c.Website, c.Phone)
	return company.CreateCompany()
}

func (c Company) DeleteCompany() error {
	company := repository.NewCompany(c.ID, c.Name, c.Code, c.Country, c.Website, c.Phone)
	err := company.DeleteCompany()

	if err != nil {
		return err
	}

	produce := Produce{
		Event:     "EVENT_DELETE",
		CompanyId: strconv.Itoa(c.ID),
	}

	go produce.SendMutationMessage()

	return nil
}

func (c Company) UpdateCompany(id string) error {
	company := repository.NewCompany(c.ID, c.Name, c.Code, c.Country, c.Website, c.Phone)
	err := company.UpdateCompany(id)

	if err != nil {
		return err
	}

	produce := Produce{
		Event:     "EVENT_UPDATE",
		CompanyId: id,
	}

	go produce.SendMutationMessage()

	return err
}
