package common

import (
	"github.com/tekab-dev/tekab-test/models"
	"strconv"
	"strings"
)

func ChoiceIndex(tab []models.QuestionChoices, id uint64) int {
	for i := 0; i < len(tab); i++ {
		if tab[i].ChoiceID == id {
			return i
		}
	}
	return -1
}
func GetTestCandidate(idTestcandidat string) (uint64, error, uint64, error) {
	idTest, errIdTest := strconv.ParseUint(idTestcandidat[strings.Index(idTestcandidat, "=")+1:strings.Index(idTestcandidat, "&")], 10, 64)
	idCandidate, errIdCandidate := strconv.ParseUint(idTestcandidat[strings.LastIndex(idTestcandidat, "=")+1:], 10, 64)
	return idTest, errIdTest, idCandidate, errIdCandidate
}
