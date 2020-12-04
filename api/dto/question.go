package dto

import (
	"errors"
	"github.com/nerock/questionapi/domain"
	"time"
)

const (
	emptyQuestionTextError = "question must have text"
	timeFormat             = "2006-01-02 15:04:05"
)

type NewQuestionRequest struct {
	Text    string   `json:"text"`
	Choices []Choice `json:"choices"`
}

type QuestionResponse struct {
	Text      string   `json:"text"`
	CreatedAt string   `json:"createdAt"`
	Choices   []Choice `json:"choices"`
}

func MapToDomainQuestion(request NewQuestionRequest) (domain.Question, error) {
	if len(request.Text) == 0 {
		return domain.Question{}, errors.New(emptyQuestionTextError)
	}

	choices, err := mapToDomainChoices(request.Choices)
	if err != nil {
		return domain.Question{}, err
	}

	return domain.Question{
		Text:      request.Text,
		CreatedAt: time.Now(),
		Choices:   choices,
	}, nil
}

func MapToQuestionsReponse(questions []domain.Question) []QuestionResponse {
	questionsResponse := make([]QuestionResponse, 0, len(questions))

	for _, q := range questions {
		questionsResponse = append(questionsResponse, MapToQuestionReponse(q))
	}

	return questionsResponse
}

func MapToQuestionReponse(question domain.Question) QuestionResponse {
	return QuestionResponse{
		Text:      question.Text,
		CreatedAt: question.CreatedAt.Format(timeFormat),
		Choices:   mapToDtoChoices(question.Choices),
	}
}
