package validators

import (
	"GoAuth2/main/error_models"
	"net/http"
)

func IsInvalidUsername(username string) *error_models.ApplicationError {
	if username == "" {
		return &error_models.ApplicationError{
			Message:    "invalid username",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
	}
	return nil
}
