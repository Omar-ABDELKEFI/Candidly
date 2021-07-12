package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func GetUser(email string) models.User {
	log.Println("getting user from db ...")
	db, _ := database.GetDb()
	var user models.User
	db.First(&user, "email = ?", email)
	return user
}
