package handler

import (
	"encoding/json"
	"github.com/raymondgitonga/company-service/internal/db"
	"net/http"
)

type CompaniesResponse struct {
	Message   string       `json:"message"`
	Companies []db.Company `json:"companies"`
}

type CompanyResponse struct {
	Message string     `json:"message"`
	Company db.Company `json:"company"`
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	company := buildCompany(r)

	res, err := company.GetCompanies()

	companyResponse := CompaniesResponse{
		Message:   "success",
		Companies: res,
	}

	if err != nil {
		companyResponse.Message = err.Error()
		jsonResponse, _ := json.Marshal(companyResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}

	jsonResponse, _ := json.Marshal(companyResponse)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	company := buildCompany(r)

	res, err := company.GetCompany()

	companyResponse := CompanyResponse{
		Message: "success",
		Company: res,
	}

	if err != nil {
		companyResponse.Message = err.Error()
		jsonResponse, _ := json.Marshal(companyResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}

	jsonResponse, _ := json.Marshal(companyResponse)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func buildCompany(r *http.Request) db.ICompany {
	name := r.URL.Query().Get("name")
	code := r.URL.Query().Get("code")
	country := r.URL.Query().Get("country")
	website := r.URL.Query().Get("website")
	phone := r.URL.Query().Get("phone")

	company := db.NewCompany(name, code, country, website, phone)

	return company
}
