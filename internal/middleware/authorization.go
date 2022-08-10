package middleware

import (
	"encoding/json"
	"errors"
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

	person, err := db.NewPerson(email).GetPerson()

	fmt.Println(person)

	if err != nil || len(person.Email) <= 0 {
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
	tokenString, err := generate(email, person.Role)

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

	authorization := Authorization{
		Message:    "success",
		Expiration: time.Now().Add(time.Minute * 43200),
		Email:      email,
		Token:      tokenString,
	}

	jsonResponse, _ := json.Marshal(authorization)

	fmt.Println(IsAuthorized(tokenString))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}

// Authorize checks if a request is authorized
func Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		isAuth, err := IsAuthorized(token)
		if err != nil || !isAuth {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(err.Error()))
			return
		}
		handler(w, r)
	}
}

func IsAuthorized(tokenString string) (bool, error) {
	var keyfunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	}

	parsed, err := jwt.Parse(tokenString, keyfunc)
	if err != nil {
		return false, err
	}

	if !parsed.Valid {
		return false, err
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		if claims["role"] != "admin" {
			return false, errors.New("operation not allowed for user")
		}
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("error validating token: %s", err.Error()))
	}
}

// generate generates a jwt token
func generate(email string, role string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["role"] = role

	return token.SignedString(mySigningKey)
}
