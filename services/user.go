package services

import "github.com/tekab-dev/tekab-test/repositories"

func ValidLogin(email string, password string) (valid bool) {
	user := repositories.GetUser()
	if user.Email == email && user.Password == password {
		return true
	}
	return false
}
