package utils

import (
	"github.com/tonygits/test-lending-svc/entities"
)

func GenerateUser(username string, userLimit float64) (user *entities.User) {
	user = &entities.User{
		Username:    username,
		UserLimit:   userLimit,
	}

	return user
}
