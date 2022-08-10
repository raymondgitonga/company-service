package middleware

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	email := "raymond@test.com"
	role := "admin"

	token, err := generate(email, role)
	assert.NoError(t, err)
	assert.NotNil(t, token)
}

func TestIsAuthorized(t *testing.T) {
	type testCase struct {
		name     string
		email    string
		role     string
		expected bool
	}

	testCases := []testCase{
		{
			name:     "role admin",
			email:    "raymond@test.com",
			role:     "admin",
			expected: true,
		},
		{
			name:     "role user",
			email:    "gitonga@test.com",
			role:     "user",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, _ := generate(tc.email, tc.role)

			isAuth, _ := IsAuthorized(token)
			assert.Equal(t, tc.expected, isAuth)
		})
	}
}
