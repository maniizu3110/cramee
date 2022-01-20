package main

import (
	"cramee/api"
	"cramee/util"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logLevel)
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db := util.InitDatabase(config)
	server, err := api.NewServer(db, config)
	if err != nil {
		log.Fatal("cannot create new server:", err)
	}

	server.Start(config.ServerAddress)
}
