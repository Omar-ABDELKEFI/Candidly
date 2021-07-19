package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func GetUser(email string) models.User {
	log.Println("getting user from db ...")
	db := database.DB
	var user models.User
	db.First(&user, "email = ?", email)
	log.Println(user.Password)
	return user
}
