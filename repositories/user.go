package repositories

import (
	"github.com/dev/test/database"
	"github.com/dev/test/models"
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
