package main

import (
	"go-supervise/db"
	"go-supervise/handlers"
	"go-supervise/server"
	"go-supervise/services"
	"go-supervise/services/checkup"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server    *server.Config
	Services  *services.Config
	Datastore *db.Config
}

func NewConfigFromYML(ymlFilePath string) *Config {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(ymlFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func main() {
	godotenv.Load()
	var config *Config
	useEnvConfig := os.Getenv("USE_ENV_CONFIG")
	ymlConfigFile := os.Getenv("CONFIG_FILE")
	switch true {
	case useEnvConfig == "" && ymlConfigFile != "":
		config = NewConfigFromYML(ymlConfigFile)
	default:
		config = NewConfigFromYML("server.config.yml")
	}

	db, err := db.NewDB(*config.Datastore)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(config.Server).Build()
	if err := handlers.NewHandlers(s).Build(); err != nil {
		log.Fatal(err)
	}

	g := errgroup.Group{}
	g.Go(func() error {
		return checkup.GetCheckUpService().RunWithInterval(
			config.Services.CheckUpService.Interval*time.Second,
			&http.Client{},
			db,
		)
	})
	g.Go(func() error {
		return s.Run()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
