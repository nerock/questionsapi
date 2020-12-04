package json

import (
	"encoding/json"
	"github.com/nerock/questionapi/domain"
	"io/ioutil"
	"log"
)

type questionsReader struct {
	source    string
	questions []domain.Question
}

func NewJsonQuestionsReader(source string) (questionsReader, error) {
	qr := questionsReader{
		source:    source,
		questions: make([]domain.Question, 0),
	}

	if err := qr.load(source); err != nil {
		return qr, err
	}

	return qr, nil
}

func (qr questionsReader) GetQuestions() []domain.Question {
	return qr.questions
}

func (qr *questionsReader) load(source string) error {
	var questions []domain.Question

	file, err := ioutil.ReadFile(source)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(file, &questions)
	if err != nil {
		log.Println(err)
		return err
	}

	qr.questions = questions

	return nil
}
