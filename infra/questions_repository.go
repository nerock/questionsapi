package infra

import (
	"errors"
	"github.com/nerock/questionapi/domain"
	"github.com/nerock/questionapi/infra/json"
	"strings"
)

const (
	JSON      = "json"
	CSV       = "csv"
	NO_READER = "wrong read source"
)

type QuestionRepository interface {
	GetQuestions() []domain.Question
	AddQuestion(question domain.Question) error
}

func GetRepository(readerType, source string) (QuestionRepository, error) {
	if strings.EqualFold(readerType, JSON) {
		return json.NewJsonQuestionReader(source)
	}

	return nil, errors.New(NO_READER)
}
