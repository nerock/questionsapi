package api

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nerock/questionapi/domain"
	"github.com/nerock/questionapi/infra"
)

const (
	REQUIRED_LANG = "valid ISO 639-1 language code is required as parameter with 'lang' key"
	ISO639_1      = 2
)

type server struct {
	repo infra.QuestionRepository
	e    *echo.Echo
}

func NewServer(repository infra.QuestionRepository) server {
	e := echo.New()
	e.Use(middleware.Logger())

	return server{
		repo: repository,
		e:    e,
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
		return c.JSON(400, errors.New(REQUIRED_LANG))
	}

	return c.JSON(200, s.repo.GetQuestions())
}

func (s server) addQuestion(c echo.Context) error {
	var question domain.Question
	if err := c.Bind(&question); err != nil {
		return err
	}

	s.repo.AddQuestion(question)

	return c.JSON(200, nil)
}
