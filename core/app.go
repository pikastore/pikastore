package core

// Defines the version whatnot 
type App struct {
	version string
}

var app = App{
	version: "0",
}

type Config struct {
	Data  string
	DbCon string
}

//Defaults
const (
	DefaultData  string = "/ps_data"
	DefaultDbcon string = "/ps_db"
)
//Creats the App
func NewApp(config Config) *Config {
	if config.DbCon == "" {
		config.DbCon = DefaultDbcon
	}
	if config.Data == "" {
		config.Data = DefaultData
	}
	return &Config{
		Data:  config.Data,
		DbCon: config.DbCon,
	}
}
