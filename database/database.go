package database

import (
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func GetDb() {
	const DNS = "root:dnVh9M9g3Q@tcp(mysql_rest:3306)/tekabTest?parseTime=true"
	log.Println("before connection ..")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Println("Cannot connect to Database")
		return
	}
	DB = db
	log.Println("Connected to database")

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
	db.SetupJoinTable(&models.Candidate{}, "Test", &models.TestCandidate{})
	db.AutoMigrate(
		&models.User{},
		&models.Skill{},
		&models.Question{},
		&models.Choices{},
		&models.Test{},
		&models.Candidate{},
		&models.TestQuestion{},
		&models.Answer{},
		&models.AnswerChoices{},
	)

}
