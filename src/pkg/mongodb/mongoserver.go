package mongodb

import (
	"studybud/src/cmd/utils"
	"studybud/src/pkg/client/mongodb"
	"studybud/src/pkg/repository"

	"github.com/sirupsen/logrus"
)

func Initialize(config utils.Configuration, collection string) repository.MongoRepository {
	var log = logrus.New()

	log.WithFields(logrus.Fields{
		"mongo_url":   config.Database.Url,
		"server_port": config.Server.Port,
		"db_name":     config.Database.DbName,
		"collection":  collection,
		"timeout":     config.App.Timeout,
	}).Info("\nConfiguration information\n")

	logrus.Infof("db connection has started")

	client, err := mongodb.ConnectMongoDb(config.Database.Url)

	config.Database.Collection = collection
	if err != nil {
		logrus.Fatal(err)
	}

	repository := repository.NewMongoRepository(&config, client)

	return repository
}
