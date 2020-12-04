package api

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nerock/questionapi/api/dto"
	"github.com/nerock/questionapi/infra"
)

const (
	RequiredLang = "valid ISO 639-1 language code is required as parameter with 'lang' key"
	ISO639_1     = 2
)

type server struct {
	repo       infra.QuestionRepository
	translator infra.Translator
	e          *echo.Echo
}

func NewServer(repository infra.QuestionRepository, translator infra.Translator) server {
	e := echo.New()
	e.Use(middleware.Logger())

	return server{
		repo:       repository,
		translator: translator,
		e:          e,
	}
}

func (s server) Run() {
	s.routes()
	s.e.Logger.Fatal(s.e.Start(":8080"))
}

func (s server) routes() {
	s.e.GET("/questions", s.getQuestions)
	s.e.POST("/questions", s.addQuestion)
}

func (s server) getQuestions(c echo.Context) error {
	lang := c.QueryParam("lang")
	if len(lang) != ISO639_1 {
		return c.JSON(400, errors.New(RequiredLang))
	}

	questions := s.repo.GetQuestions()
	translatedQuestions, err := s.translator.TranslateQuestions(questions, lang)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, dto.MapToQuestionsReponse(translatedQuestions))
}

func (s server) addQuestion(c echo.Context) error {
	var questionRequest dto.NewQuestionRequest
	if err := c.Bind(&questionRequest); err != nil {
		return err
	}

	question, err := dto.MapToDomainQuestion(questionRequest)
	if err != nil {
		return c.JSON(400, err)
	}

	addedQuestion, err := s.repo.AddQuestion(question)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, dto.MapToQuestionReponse(addedQuestion))
}
