package validators

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIsInvalidUsernameNotValidUsername(t *testing.T) {
	err := IsInvalidUsername("")

	assert.NotNil(t, err, "expected error is nil")
	assert.EqualValues(t, "invalid username", err.Message)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "bad_request", err.Code)
}

func TestIsInvalidUsernameNoError(t *testing.T) {
	err := IsInvalidUsername("validUsername")
	assert.Nil(t, err, "expected error is not nil")
}
