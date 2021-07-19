package database

import (
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GetDb() (*gorm.DB, error) {
	const DNS = "root:dnVh9M9g3Q@tcp(mysql_rest:3306)/tekabTest?parseTime=true"
	log.Println("before connection ..")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Println("Cannot connect to Database")
		return nil, err
	}
	log.Println("Connected to database")

	return db, err
}
func MigrateDatabase() {
	const DNS = "root:dnVh9M9g3Q@tcp(mysql_rest:3306)/tekabTest?parseTime=true"
	log.Println("before connection ..")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Cannot connect to Database")

	}
	log.Println("Connected to database")
	db.AutoMigrate(
		&models.User{},
		&models.Skill{},
		&models.Question{},
		&models.Choices{},
		&models.Test{},
		&models.Candidate{},
		&models.TestQuestion{},
		&models.Answer{},
		&models.TestCandidate{},
		&models.AnswerChoices{},
	)

}
