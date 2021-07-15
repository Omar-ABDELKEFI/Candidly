package common

import "github.com/tekab-dev/tekab-test/models"

func ChoiceIndex(tab []models.AnswerDbResponse, id uint64) int {
	for i := 0; i < len(tab); i++ {
		if tab[i].ChoiceId == id {
			return i
		}
	}
	return -1
}
