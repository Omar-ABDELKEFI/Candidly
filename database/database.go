package database

import (
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func GetDb() {
	const DNS = "root:dnVh9M9g3Q@tcp(mysql_rest:3306)/tekabTest?parseTime=true"
	log.Println("before connection ..")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err.Error())
		log.Println("Cannot connect to Database")
		return
	}

	DB = db
	log.Println("Connected to database")
	//migration
	MigrateDatabase(db)

}
func MigrateDatabase(db *gorm.DB) {
	log.Println("before migration ..")

	db.SetupJoinTable(&models.Candidate{}, "Test", &models.TestCandidate{})
	db.SetupJoinTable(&models.Test{}, "Candidate", &models.TestCandidate{})
	db.SetupJoinTable(&models.Question{}, "Test", &models.TestQuestion{})
	db.SetupJoinTable(&models.Test{}, "Question", &models.TestQuestion{})
	db.AutoMigrate(
		&models.User{},
		&models.Skill{},
		&models.Question{},
		&models.Choices{},
		&models.Test{},
		&models.Candidate{},
		&models.Answer{},
		&models.AnswerChoices{},
	)
	log.Println("after to database")
}
