package json

import (
	"github.com/nerock/questionapi/domain"
)

type ChoiceVO struct {
	Text string `json:"text"`
}

func mapToQuestionChoices(choicesRequest []ChoiceVO) []domain.Choice {
	choices := make([]domain.Choice, 0)
	for _, c := range choicesRequest {
		choices = append(choices, mapToQuestionChoice(c))
	}

	return choices
}

func mapToQuestionChoice(choicesRequest ChoiceVO) domain.Choice {
	return domain.Choice{Text: choicesRequest.Text}
}
