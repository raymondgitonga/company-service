package handler

import (
	"encoding/json"
	"github.com/raymondgitonga/company-service/internal/repository"
	"github.com/raymondgitonga/company-service/internal/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CompaniesResponse struct {
	Message   string               `json:"message"`
	Companies []repository.Company `json:"companies"`
}

type CompanyResponse struct {
	Message string              `json:"message"`
	Company *repository.Company `json:"company"`
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	company := buildCompanyStruct(r)

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
	company := buildCompanyStruct(r)

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
	req := repository.Company{}

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

	company := repository.NewCompany(req.ID, req.Name, req.Code, req.Country, req.Website, req.Phone)

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

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	companyId, _ := strconv.Atoi(r.URL.Query().Get("id"))

	company := service.NewCompany(companyId, "", "", "", "", "")

	err := company.DeleteCompany()

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

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	companyId := r.URL.Query().Get("id")
	req := repository.Company{}

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

	company := repository.NewCompany(0, req.Name, req.Code, req.Country, req.Website, req.Phone)

	err = company.UpdateCompany(companyId)

	if err != nil {
		jsonResponse, _ := json.Marshal("message: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	produce := service.Produce{
		Event:     "EVENT_UPDATE",
		CompanyId: companyId,
	}

	// Sends update message to kafka topic async
	go produce.SendMutationMessage()

	jsonResponse, _ := json.Marshal("message: success")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func buildCompanyStruct(r *http.Request) service.CompanyService {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	name := r.URL.Query().Get("name")
	code := r.URL.Query().Get("code")
	country := r.URL.Query().Get("country")
	website := r.URL.Query().Get("website")
	phone := r.URL.Query().Get("phone")

	company := service.NewCompany(id, name, code, country, website, phone)

	return company
}
