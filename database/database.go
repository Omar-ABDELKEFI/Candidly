package database

import (
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GetDb() (*gorm.DB, error) {
	const DNS = "root:dnVh9M9g3Q@tcp(mysql_rest:3306)/tekabTest"
	log.Println("before connection ..")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Cannot connect to Database")
		return nil, err
	}
	log.Println("Connected to database")
	MigrateDatabase(db)
	return db, err
}
func MigrateDatabase(db *gorm.DB) {

	db.AutoMigrate(
		&models.Skill{},
		&models.Question{},
		&models.Text{},
		&models.MultipleChoiceQuestion{},
		&models.Choices{},
		&models.Test{},
		&models.NotifyEmails{},
		&models.Candidat{},
		&models.Note{},
	)

}
