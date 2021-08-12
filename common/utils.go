package common

import "github.com/tekab-dev/tekab-test/models"

func ChoiceIndex(tab []models.QuestionChoices, id uint64) int {
	for i := 0; i < len(tab); i++ {
		if tab[i].ChoiceID == id {
			return i
		}
	}
	return -1
}
