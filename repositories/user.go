package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	_ "golang.org/x/crypto/bcrypt"
	_ "log"
)

func GetUser(email string) models.User {
	/*hashedPass, err := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
	if err != nil {

		log.Println("\"unable to hash password\"", err)

		return
	}*/
	db, _ := database.GetDb()
	var user models.User
	db.First(&user, "email = ?", email)
	return user
}
