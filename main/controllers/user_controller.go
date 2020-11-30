package controllers

import (
	"GoAuth2/main/domain"
	"GoAuth2/main/error_models"
	"GoAuth2/main/services"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"

	userValidator "GoAuth2/main/validators"
)

func GetUserByUsername(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	if isInvalidUsername(response, username) {
		return
	}
	user, done := getUserFromService(response, username)
	if done {
		return
	}
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}

func isInvalidUsername(response http.ResponseWriter, username string) bool {
	usernameValidationError := userValidator.IsInvalidUsername(username)
	if usernameValidationError != nil {
		fillErrorResponse(response, usernameValidationError)
		return true
	}
	return false
}

func getUserFromService(response http.ResponseWriter, username string) (*domain.User, bool) {
	user, err := services.GetUser(username)
	if err != nil {
		fillErrorResponse(response, err)
		return nil, true
	}
	return user, false
}

func fillErrorResponse(response http.ResponseWriter, err *error_models.ApplicationError) {
	userNotFoundErrorJson, _ := json.Marshal(err)
	response.WriteHeader(err.StatusCode)
	response.Write(userNotFoundErrorJson)
}
