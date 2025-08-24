package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	return &Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetInt("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
	}
}
