package repositories

import (
	"github.com/tekab-dev/tekab-test/models"
)

func GetUser() (user models.User) {
	User := models.User{"fff", "adminadmin"}
	return User
}
