package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tekab-dev/tekab-test/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func ValidateLogin(email string, password string) (string, error) {
	user := repositories.GetUser(email)
	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	log.Println(errHash, "error")
	var err error
	if user.Email != email || errHash != nil {
		err = errors.New("Bad Credentials")
		return "", err
	}
	token, err := CreateToken(email)
	if err != nil {
		return "", err
	}
	return token, err

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
