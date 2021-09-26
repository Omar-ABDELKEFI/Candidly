package repositories

import (
	"fmt"
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateCandidate(candidate []models.Candidate) ([]models.Candidate, error) {

	db := database.DB

	fmt.Println("created candidate : ", candidate)
	log.Println(candidate, "candidateCandidateCandidate\"")
	err := db.Create(&candidate).Error

	if err != nil {
		return candidate, err
	}
	for _, s := range candidate {
		fmt.Printf("%+v\n", s.Email)
		log.Println("oooo")
	}
	return candidate, nil
}
func SelectCandidateEmails() ([]string, error) {
	log.Println("Creating Candidate ...")
	var emails []string
	db := database.DB
	err := db.Table("candidates").Select("email").Scan(&emails).Error
	log.Println(emails, "emailsemails")
	if err != nil {
		return emails, err
	}

	return emails, nil
}
func FindCandidate(emails []string) ([]models.Candidate, error) {
	var candidatesId []models.Candidate
	db := database.DB
	err := db.Table("candidates").Where("email In ?", emails).Find(&candidatesId).Error
	log.Println(emails, "aaaa")
	if err != nil {
		return candidatesId, err
	}

	for _, s := range candidatesId {
		fmt.Printf("%+v\n candidatesId", s)
	}
	return candidatesId, nil
}
