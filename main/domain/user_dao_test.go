package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser("abc")

	assert.Nil(t, user, "expected user is not nil")
	assert.NotNil(t, err, "expected error is nil")
	assert.NotNil(t, err.Message, "expected error message")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser("steve")

	assert.NotNil(t, user, "expected user is nil")
	assert.Nil(t, err, "expected error is not nil")
	assert.EqualValues(t, "steve", user.Username)
	assert.NotNil(t, user.Email, "users must have email")
	assert.NotNil(t, user.Password, "users must have password")
}
