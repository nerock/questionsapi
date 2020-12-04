package main

import (
	"github.com/nerock/questionapi/api"
	"github.com/nerock/questionapi/infra"
)

func main() {
	r, err := infra.GetReader("json", "questions.json")
	if err != nil {
		panic(err)
	}

	srv := api.NewServer(r)
	srv.Run()
}
