package utils

import (
	"encoding/base64"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Read_Configuration(config Configuration) Configuration {

	mongoUri := os.Getenv("MONGODB_URL")
	port := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	collection := os.Getenv("COLLECTION")
	appName := os.Getenv("APP_NAME")
	requestTimeOut := os.Getenv("TIMEOUT")

	if mongoUri != "" && port != "" && dbName != "" && collection != "" && appName != "" {
		requestTimeOut, err := strconv.Atoi(requestTimeOut)
		if err != nil {
			logrus.Fatalf("TIMEOUT cannot be converted %d\n", requestTimeOut)
		}
		logrus.Info("timeout data type", reflect.TypeOf(requestTimeOut))

		return Configuration{
			App:      Application{Name: appName, Timeout: config.App.Timeout},
			Database: DatabaseSettings{Url: mongoUri, DbName: dbName, Collection: collection},
			Server:   ServerSettings{Port: port},
		}
	}

	return Configuration{
		App:      Application{Name: config.App.Name, Timeout: config.App.Timeout},
		Database: DatabaseSettings{Url: config.Database.Url, DbName: config.Database.DbName, Collection: config.Database.Collection},
		Server:   ServerSettings{Port: config.Server.Port},
	}
}

func Read() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		logrus.Errorf("unable to decode into struct, %v", err)
	}

	logrus.Warnf("config with variables %v", config)

	return config
}

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		logrus.Fatalf("failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
