package json

import (
	"encoding/json"
	"github.com/nerock/questionapi/domain"
	"io/ioutil"
)

type questionRepository struct {
	source    string
	questions []domain.Question
}

func NewJsonQuestionRepository(source string) (*questionRepository, error) {
	qr := questionRepository{
		source:    source,
		questions: make([]domain.Question, 0),
	}

	if err := qr.load(source); err != nil {
		return nil, err
	}

	return &qr, nil
}

func (qr questionRepository) GetQuestions() []domain.Question {
	return qr.questions
}

func (qr *questionRepository) AddQuestion(question domain.Question) (domain.Question, error) {
	qr.questions = append(qr.questions, question)

	return question, nil
}

func (qr *questionRepository) load(source string) error {
	var questions []QuestionVO

	file, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &questions)
	if err != nil {
		return err
	}

	qr.questions = MapToQuestions(questions)

	return nil
}
