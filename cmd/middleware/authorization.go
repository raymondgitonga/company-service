package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/raymondgitonga/company-service/internal/db"
	"net/http"
	"os"
	"time"
)

type Authorization struct {
	Message    string    `json:"message"`
	Expiration time.Time `json:"Expiration"`
	Email      string    `json:"Email"`
	Token      string    `json:"Token"`
}

func GenerateJWT(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	secretKey := os.Getenv("JWT_SECRET_KEY")
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err := fmt.Sprintf("something went wrong: %s", err.Error())

		response := Authorization{
			Message:    err,
			Expiration: time.Time{},
			Email:      "",
			Token:      "",
		}

		jsonResponse, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(jsonResponse)
		return
	}

	exists, err := db.NewPerson(email).PersonExists()

	if !exists || err != nil {
		authorization := Authorization{
			Message:    "user does not exist",
			Expiration: time.Time{},
			Email:      email,
			Token:      "",
		}

		jsonResponse, _ := json.Marshal(authorization)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(jsonResponse)
		return
	}

	authorization := Authorization{
		Message:    "success",
		Expiration: time.Now().Add(time.Minute * 30),
		Email:      email,
		Token:      tokenString,
	}

	jsonResponse, _ := json.Marshal(authorization)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}
