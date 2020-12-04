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

type QuestionsReader interface {
	GetQuestions() []domain.Question
}

func GetReader(readerType, source string) (QuestionsReader, error) {
	if strings.EqualFold(readerType, JSON) {
		return json.NewJsonQuestionsReader(source)
	}

	return nil, errors.New(NO_READER)
}
