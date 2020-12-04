package main

import (
	"github.com/nerock/questionapi/api"
	"github.com/nerock/questionapi/infra"
	"github.com/spf13/viper"
)

const (
	sourceKey     = "source"
	repositoryKey = "repository"
	googleKey     = "googleKey"
)

func main() {
	readConfig()

	t, err := infra.GetTranslator(viper.GetString(googleKey))
	if err != nil {
		panic(err)
	}

	r, err := infra.GetRepository(viper.GetString(repositoryKey), viper.GetString(sourceKey))
	if err != nil {
		panic(err)
	}

	srv := api.NewServer(r, t)
	srv.Run()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Could not read configuration")
	}
}
