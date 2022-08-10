package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckIpAddressIsAllowed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "https://ipapi.co/104.28.60.46/json/")
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	api := API{server.Client(), "https://ipapi.co/104.28.60.46/json/"}

	isAllowed := api.checkIpAddressIsAllowed()

	assert.Equal(t, true, isAllowed)
}

func TestCheckIpAddressNotAllowed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "https://ipapi.co/197.232.130.34/json/")
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	api := API{server.Client(), "https://ipapi.co/197.232.130.34/json/"}

	isAllowed := api.checkIpAddressIsAllowed()

	assert.Equal(t, false, isAllowed)
}
