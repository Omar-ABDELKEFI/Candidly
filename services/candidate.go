package services

import (
	"github.com/tekab-dev/tekab-test/common"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

func CreateCandidate(candidates []models.Candidate) ([]models.TestCandidate, []models.Candidate, []string, []string, error) {
	//Create new candidate
	var testCandidates []models.TestCandidate
	var testCandidate models.TestCandidate

	emails, err := repositories.SelectCandidateEmails()
	var newCan []models.Candidate
	var oldEmail []string
	for _, candidate := range candidates {

		if !common.Contains(emails, candidate.Email) {
			newCan = append(newCan, candidate)
		} else {
			oldEmail = append(oldEmail, candidate.Email)
		}
	}
	newCandidates, err := repositories.CreateCandidate(newCan)
	if err != nil && len(newCandidates) != 0 {
		return nil, candidates, []string{}, []string{}, err
	}
	candidatesId, err := repositories.FindCandidate(oldEmail)
	if err != nil && len(candidatesId) != 0 {
		return nil, candidates, []string{}, []string{}, err
	}
	for _, id := range candidatesId {

		testCandidate.CandidateID = uint64(id.ID)
		testCandidate.TestID = uint64(candidates[0].Test[0].ID)
		testCandidates = append(testCandidates, testCandidate)
	}
	allCandidate := append(newCandidates, candidatesId...)
	log.Println(allCandidate, "allCandidate")
	emailsDuplicate, newTestCandidate, errDuplicate := repositories.CreateTestCandidate(testCandidates)

	log.Println(len(errDuplicate) != 0, "len(errDuplicate) != 0")
	if len(errDuplicate) != 0 && len(candidatesId) != 0 {
		var allCandidateNoteDuplicate []models.Candidate
		for _, can := range allCandidate {
			log.Println("testAllCandidate")
			if !common.Contains(emailsDuplicate, can.Email) {
				allCandidateNoteDuplicate = append(allCandidateNoteDuplicate, can)
			}
		}

		if len(allCandidateNoteDuplicate) > 0 {

			common.Send(allCandidateNoteDuplicate, candidates[0].Test[0].ID)
		}

		return newTestCandidate, candidates, emailsDuplicate, errDuplicate, nil
	}

	common.Send(allCandidate, candidates[0].Test[0].ID)

	return newTestCandidate, candidates, emailsDuplicate, nil, nil
}
