package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"studybud/src/cmd/utils"
	"studybud/src/pkg/handlers"
	"studybud/src/pkg/mongodb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var router *chi.Mux

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/favicon.ico")
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	config := read_configuration(read())

	mongodb.Initialize(config)

	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error
	handlers.Catch(err)

	handlers.LoadTemplates()

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/", handlers.HomeHandler)

	port := ":8080"
	fmt.Println("server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func read_configuration(config utils.Configuration) utils.Configuration {

	mongoUri := os.Getenv("MONGODB_URL")
	port := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	collection := os.Getenv("COLLECTION")
	appName := os.Getenv("APP_NAME")
	requestTimeOut := os.Getenv("TIMEOUT")

	if mongoUri != "" && port != "" && dbName != "" && collection != "" && appName != "" {
		requestTimeOut, err := strconv.Atoi(requestTimeOut)
		if err != nil {
			logrus.Fatal("TIMEOUT cannot be converted %d\n", requestTimeOut)
		}
		logrus.Info("timeout data type", reflect.TypeOf(requestTimeOut))

		return utils.Configuration{
			App:      utils.Application{Name: appName, Timeout: config.App.Timeout},
			Database: utils.DatabaseSettings{Url: mongoUri, DbName: dbName, Collection: collection},
			Server:   utils.ServerSettings{Port: port},
		}
	}

	return utils.Configuration{
		App:      utils.Application{Name: config.App.Name, Timeout: config.App.Timeout},
		Database: utils.DatabaseSettings{Url: config.Database.Url, DbName: config.Database.DbName, Collection: config.Database.Collection},
		Server:   utils.ServerSettings{Port: config.Server.Port},
	}
}

func read() utils.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		logrus.Error("unable to decode into struct, %v", err)
	}

	logrus.Warn("config with variables %v", config)

	return config
}
