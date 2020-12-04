package infra

import (
	"errors"
	"github.com/nerock/questionapi/domain"
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

func GetRepository(readerType, source string) (QuestionRepository, error) {
	if strings.EqualFold(readerType, jsonExt) {
		return json.NewJsonQuestionRepository(source)
	}

	return nil, errors.New(noRepository)
}
