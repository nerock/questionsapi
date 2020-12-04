package csv

import (
	"encoding/csv"
	"errors"
	"github.com/nerock/questionapi/domain"
	"io"
	"os"
	"time"
)

const (
	noQuestionsError = "there are no questions in this csv file"
	fieldNumber      = 5
	timeFormat       = "2006-01-02 15:04:05"
)

type questionRepository struct {
	source    string
	questions []domain.Question
}

func NewCsvQuestionRepository(source string) (*questionRepository, error) {
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
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(file)
	_, err = csvReader.Read() // Skip fields names line
	if err == io.EOF {
		return errors.New(noQuestionsError)
	}

	lines, err := csvReader.ReadAll()

	qr.questions = csvLinesToQuestions(lines)
	return nil
}

func csvLinesToQuestions(lines [][]string) []domain.Question {
	questions := make([]domain.Question, 0, len(lines))

	for _, line := range lines {
		if len(line) != fieldNumber {
			continue
		}

		question, err := csvLineToQuestion(line)
		if err == nil {
			questions = append(questions, question)
		}
	}

	return questions
}

func csvLineToQuestion(line []string) (domain.Question, error) {
	createdAt, err := time.Parse(timeFormat, line[1])
	if err != nil {
		return domain.Question{}, err
	}

	return domain.Question{
		Text:      line[0],
		CreatedAt: createdAt,
		Choices: []domain.Choice{
			{Text: line[2]},
			{Text: line[3]},
			{Text: line[4]},
		},
	}, nil
}
