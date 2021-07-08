package repositories

import (
	"github.com/tekab-dev/tekab-test/models"
)

func GetUser() (user models.User) {
	User := models.User{"1zzz", "adminadmin"}
	return User
}
