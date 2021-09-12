package common

import (
	"github.com/go-playground/validator"
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

type ErrorResponse struct {
	StructField string
	Tag         string
	Field       string
	Param       string
	Value       interface{}
}

func ValidateStruct(str interface{}) ([]*ErrorResponse, error) {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(str)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.StructField = err.StructField()
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Param = err.Param()
			element.Value = err.Value()
			errors = append(errors, &element)
		}
	}
	return errors, err
}
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
