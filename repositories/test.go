package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTest(test models.Test) (models.Test, error) {
	log.Println("Creating test ...")
	db, err := database.GetDb()
	if err != nil {
		return test, err
	}
	err = db.Create(&test).Error

	if err != nil {
		return test, err
	}
	log.Println("created test : ", test)

	return test, nil
}
