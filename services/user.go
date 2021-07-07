package services

import "github.com/tekab-dev/tekab-test/repositories"

func ValidLogin(emial string, password string) (valid bool) {
	if repositories.ReturnUser().Email == emial && repositories.ReturnUser().Password == password {
		return true
	}
	return false
}
