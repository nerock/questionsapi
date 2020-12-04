package json

import (
	"github.com/labstack/gommon/log"
	"github.com/nerock/questionapi/domain"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type QuestionVO struct {
	Text      string     `json:"text"`
	CreatedAt string     `json:"createdAt"`
	Choices   []ChoiceVO `json:"choices"`
}

func MapToQuestions(questionsVO []QuestionVO) []domain.Question {
	questions := make([]domain.Question, 0, len(questionsVO))
	for _, q := range questionsVO {
		question, err := MapToQuestion(q)
		if err != nil {
			log.Error(err)
		}

		questions = append(questions, question)
	}

	return questions
}

func MapToQuestion(questionVO QuestionVO) (domain.Question, error) {
	choices := mapToQuestionChoices(questionVO.Choices)

	createdAt, err := stringToTime(questionVO.CreatedAt)
	if err != nil {
		return domain.Question{}, err
	}

	return domain.Question{
		Text:      questionVO.Text,
		CreatedAt: createdAt,
		Choices:   choices,
	}, nil
}

func stringToTime(createdAt string) (time.Time, error) {
	return time.Parse(timeFormat, createdAt)
}
