package api

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nerock/questionapi/infra"
)

const (
	REQUIRED_LANG = "valid ISO 639-1 language code is required as parameter with 'lang' key"
	ISO639_1      = 2
)

type server struct {
	reader infra.QuestionRepository
	e      *echo.Echo
}

func NewServer(reader infra.QuestionRepository) server {
	e := echo.New()
	e.Use(middleware.Logger())

	return server{
		reader: reader,
		e:      e,
	}
}

func (s server) Run() {
	s.routes()
	s.e.Logger.Fatal(s.e.Start(":8080"))
}

func (s server) routes() {
	s.e.GET("/questions", s.getQuestions)
}

func (s server) getQuestions(c echo.Context) error {
	lang := c.QueryParam("lang")
	if len(lang) != ISO639_1 {
		return c.JSON(400, errors.New(REQUIRED_LANG))
	}

	return c.JSON(200, s.reader.GetQuestions())
}
