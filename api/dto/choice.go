package dto

import (
	"errors"
	"github.com/nerock/questionapi/domain"
)

const (
	emptyChoiceTextError = "question must have text"
	choiceNumberError    = "there must be exactly three choices"
	choiceNumber         = 3
)

type Choice struct {
	Text string `json:"text"`
}

func mapToDomainChoices(choicesRequest []Choice) ([]domain.Choice, error) {
	if len(choicesRequest) != choiceNumber {
		return []domain.Choice{}, errors.New(choiceNumberError)
	}

	choices := make([]domain.Choice, 0, choiceNumber)
	for _, c := range choicesRequest {
		choice, err := mapToDomainChoice(c)
		if err != nil {
			return choices, err
		}

		choices = append(choices, choice)
	}

	return choices, nil
}

func mapToDomainChoice(choicesRequest Choice) (domain.Choice, error) {
	if len(choicesRequest.Text) == 0 {
		return domain.Choice{}, errors.New(emptyChoiceTextError)
	}

	return domain.Choice{Text: choicesRequest.Text}, nil
}

func mapToDtoChoices(choices []domain.Choice) []Choice {
	dtoChoices := make([]Choice, 0, len(choices))

	for _, c := range choices {
		dtoChoices = append(dtoChoices, mapToDtoChoice(c))
	}

	return dtoChoices
}

func mapToDtoChoice(choice domain.Choice) Choice {
	return Choice{Text: choice.Text}
}
