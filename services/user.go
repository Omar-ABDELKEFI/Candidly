package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tekab-dev/tekab-test/repositories"
	"os"
	"time"
)

func ValidateLogin(email string, password string) (valid bool) {
	user := repositories.GetUser()
	if user.Email == email && user.Password == password {
		return true
	}
	return false
}
func CreateToken(userEmail string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "azert") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = userEmail
	atClaims["exp"] = time.Now().Add(time.Hour * 8760).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
