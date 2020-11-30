package domain

import (
	"GoAuth2/main/error_models"
	"fmt"
	"net/http"
)

var users = map[string]*User {
	"steve": {ID: 1, Username: "steve", Password: "ab2345wr5423ekl452jh6298u902834yh", Email: "steve@gmail.com"},
}

func GetUser(username string) (*User, *error_models.ApplicationError) {
	if user := users[username]; user != nil {
		return user, nil
	}
	return nil, &error_models.ApplicationError{
		Message:    fmt.Sprintf("user with username %s not found", username),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
