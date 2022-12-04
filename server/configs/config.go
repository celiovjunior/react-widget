package configs

import (
	"github.com/spf13/viper"
)

var cnfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// start application exec
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// load configs
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cnfg = new(config)

	cnfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cnfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cnfg.DB
}

func GetServerPort() string {
	return cnfg.API.Port
}
