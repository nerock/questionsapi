package infra

import (
	"errors"
	"github.com/nerock/questionapi/domain"
	"github.com/nerock/questionapi/infra/csv"
	"github.com/nerock/questionapi/infra/json"
	"strings"
)

const (
	jsonExt      = "json"
	csvExt       = "csv"
	noRepository = "wrong repository source"
)

type QuestionRepository interface {
	GetQuestions() []domain.Question
	AddQuestion(question domain.Question) (domain.Question, error)
}

func GetRepository(repositoryType, source string) (QuestionRepository, error) {
	if strings.EqualFold(repositoryType, jsonExt) {
		return json.NewJsonQuestionRepository(source)
	} else if strings.EqualFold(repositoryType, csvExt) {
		return csv.NewCsvQuestionRepository(source)
	}

	return nil, errors.New(noRepository)
}
