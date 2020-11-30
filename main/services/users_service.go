package services

import (
	"GoAuth2/main/domain"
	"GoAuth2/main/error_models"
)

func GetUser(username string) (*domain.User, *error_models.ApplicationError) {
	return domain.GetUser(username)
}
