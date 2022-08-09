package model

type Company struct {
	ID      int    `json:"companyId"`
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type Person struct {
	ID    int    `json:"personId"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
