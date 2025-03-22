package utils

type Configuration struct {
	Database DatabaseSettings
	Server   ServerSettings
	App      Application
}

type DatabaseSettings struct {
	Url        string
	DbName     string
	Collection string
}

type ServerSettings struct {
	Port string
}

type Application struct {
	Name    string
	Timeout int
}
