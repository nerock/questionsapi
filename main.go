package main

import (
	"github.com/nerock/questionapi/api"
	"github.com/nerock/questionapi/infra"
	"github.com/spf13/viper"
)

const (
	sourceKey     = "source"
	repositoryKey = "repository"
)

func main() {
	readConfig()

	r, err := infra.GetRepository(viper.GetString(repositoryKey), viper.GetString(sourceKey))
	if err != nil {
		panic(err)
	}

	srv := api.NewServer(r)
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
