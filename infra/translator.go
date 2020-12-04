package infra

import (
	"cloud.google.com/go/translate"
	"context"
	"errors"
	"github.com/nerock/questionapi/domain"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"io/ioutil"
)

const (
	invalidLangCode = "invalid language code"
	numEquivalence  = 4
)

type Translator struct {
	tc *translate.Client
}

func GetTranslator(jsonKeyPath string) (Translator, error) {
	jsonKey, err := ioutil.ReadFile(jsonKeyPath)
	if err != nil {
		return Translator{}, err
	}

	tc, err := translate.NewClient(context.Background(), option.WithCredentialsJSON(jsonKey))
	if err != nil {
		return Translator{}, err
	}

	return Translator{tc: tc}, nil
}

func (t Translator) TranslateQuestions(questions []domain.Question, langCode string) ([]domain.Question, error) {
	lang, err := language.Parse(langCode)
	if err != nil {
		return nil, errors.New(invalidLangCode)
	}

	translations, err := t.tc.Translate(context.Background(), mapToTranslationsArray(questions), lang, nil)
	if err != nil {
		return nil, err
	}

	return setTranslatedTexts(questions, translations), nil
}

func mapToTranslationsArray(questions []domain.Question) []string {
	translationsArray := make([]string, 0, len(questions)*numEquivalence)

	for _, q := range questions {
		translationsArray = append(translationsArray, q.Text)

		for _, c := range q.Choices {
			translationsArray = append(translationsArray, c.Text)
		}
	}

	return translationsArray
}

func setTranslatedTexts(originalQuestions []domain.Question, translationsArray []translate.Translation) []domain.Question {
	translatedQuestions := make([]domain.Question, 0, len(originalQuestions))

	for i, q := range originalQuestions {
		translationsIterator := i * numEquivalence

		translatedChoices := make([]domain.Choice, 0, len(q.Choices))
		for j, _ := range q.Choices {
			translatedChoices = append(translatedChoices, domain.Choice{Text: translationsArray[translationsIterator+j+1].Text})
		}

		translatedQuestions = append(translatedQuestions, domain.Question{
			Text:      translationsArray[translationsIterator].Text,
			CreatedAt: q.CreatedAt,
			Choices:   translatedChoices,
		})
	}

	return translatedQuestions
}
