package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type IpAddress struct {
	Ip          string `json:"ip"`
	CountryCode string `json:"country_code"`
}

type API struct {
	Client  *http.Client
	baseURL string
}

func ValidateLocation(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := API{
			Client:  &http.Client{},
			baseURL: fmt.Sprintf("https://ipapi.co/%s/json/", r.Header.Get("X-REAL-IP")),
		}

		if !api.checkIpAddressIsAllowed() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Unauthorised location"))
			return
		}
		handler(w, r)
	}
}

func (a API) checkIpAddressIsAllowed() bool {
	//var req *http.Request
	var err error
	client := a.Client
	var ipAdrr IpAddress

	url := fmt.Sprintf(a.baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	if err != nil {
		log.Print("error making call: ", err)
		return false
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Print("error making call: ", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("error reading response body ", err)
		return false
	}

	sb := string(body)

	err = json.Unmarshal([]byte(sb), &ipAdrr)
	if err != nil {
		log.Print("error unmarshalling response body ", err)
		return false
	}

	defer resp.Body.Close()

	if strings.ToUpper(ipAdrr.CountryCode) != "CY" {
		log.Print(sb)
		return false
	}
	return true
}
