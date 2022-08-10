package handler

import (
	"encoding/json"
	"github.com/raymondgitonga/company-service/internal/db"
	"io/ioutil"
	"net/http"
)

type CompaniesResponse struct {
	Message   string       `json:"message"`
	Companies []db.Company `json:"companies"`
}

type CompanyResponse struct {
	Message string      `json:"message"`
	Company *db.Company `json:"company"`
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
		Company: &res,
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

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	req := db.Company{}

	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		jsonResponse, _ := json.Marshal("message: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	err = json.Unmarshal(bodyBytes, &req)

	if err != nil {
		jsonResponse, _ := json.Marshal("message: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	company := db.NewCompany(req.Name, req.Code, req.Country, req.Website, req.Phone)

	err = company.CreateCompany()

	if err != nil {
		jsonResponse, _ := json.Marshal("message: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	jsonResponse, _ := json.Marshal("message: success")
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
