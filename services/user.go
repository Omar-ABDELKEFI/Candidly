package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dev/test/models"
	"github.com/dev/test/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

var TOKEN string

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
	TOKEN = token
	if err != nil {
		return "", err
	}
	return token, err

}
func CreateToken(userEmail string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "azert") //this should be in an env file
	expirationTime := time.Now().Add(time.Hour * 8760).Unix()
	claims := &models.Claims{
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
