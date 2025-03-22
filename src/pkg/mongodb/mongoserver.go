package mongodb

import (
	"fmt"

	"studybud/src/cmd/utils"
	"studybud/src/pkg/client/mongodb"
	"studybud/src/pkg/handlers"
	"studybud/src/pkg/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Initialize(config utils.Configuration) {
	var log = logrus.New()

	log.WithFields(logrus.Fields{
		"mongo_url":   config.Database.Url,
		"server_port": config.Server.Port,
		"db_name":     config.Database.DbName,
		"collection":  config.Database.Collection,
		"timeout":     config.App.Timeout,
	}).Info("\nConfiguration information\n")

	logrus.Infof("db connection has started")

	client, err := mongodb.ConnectMongoDb(config.Database.Url)

	if err != nil {
		logrus.Fatal(err)
	}

	repository := repository.NewMongoRepository(&config, client)
	handler := handlers.NewApiHandler(client, repository, config)

	router := gin.Default()

	api := router.Group("api/v1")
	{
		api.GET("/health", handler.Healthcheck)
	}

	formattedUrl := fmt.Sprintf(":%s", config.Server.Port)

	router.Run(formattedUrl)
}
