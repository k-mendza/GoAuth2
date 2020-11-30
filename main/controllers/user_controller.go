package controllers

import (
	"GoAuth2/main/domain"
	"GoAuth2/main/services"
	"encoding/json"
	"net/http"

	userValidator "GoAuth2/main/validators"
)

func GetUserByUsername(response http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")
	if validateUsername(response, username) {
		return
	}
	user, done := getUserFromService(response, username)
	if done {
		return
	}
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}

func getUserFromService(response http.ResponseWriter, username string) (*domain.User, bool) {
	user, err := services.GetUser(username)
	if err != nil {
		userNotFoundErrorJson, _ := json.Marshal(err)
		response.WriteHeader(err.StatusCode)
		response.Write(userNotFoundErrorJson)
		return nil, true
	}
	return user, false
}

func validateUsername(response http.ResponseWriter, username string) bool {
	usernameValidationError := userValidator.IsInvalidUsername(username)
	if usernameValidationError != nil {
		usernameErrorJson, _ := json.Marshal(usernameValidationError)
		response.WriteHeader(usernameValidationError.StatusCode)
		response.Write(usernameErrorJson)
		return true
	}
	return false
}
