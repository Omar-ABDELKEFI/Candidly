package database

import (
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GetDb() (*gorm.DB, error) {
	const DNS = "root:11@tcp(mysql_rest:3306)/tekab-test"
	log.Println("before connection ..")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Cannot connect to Database")
		return nil, err
	} else {
		log.Println("Connected to database")
	}
	MigrateDatabase(db)
	return db, err
}
func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
