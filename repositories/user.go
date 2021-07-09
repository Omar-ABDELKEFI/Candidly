package repositories

import (
	"github.com/tekab-dev/tekab-test/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetUser() (user models.User) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
	if err != nil {

		log.Println("\"unable to hash password\"", err)

		return
	}
	User := models.User{"admin@gmail.com", string(hashedPass)}
	return User
}
