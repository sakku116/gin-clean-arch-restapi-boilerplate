package config

import (
	"github.com/spf13/viper"
)

type EnvsSchema struct {
	HOST                   string
	PORT                   int
	JWT_SECRET_KEY         string
	JWT_EXP_HOURS          int
	JWT_REFRESH_SECRET_KEY string
	JWT_REFRESH_EXP_HOURS  int
	MONGO_URI              string
	MONGO_DB_NAME          string
}

var Envs *EnvsSchema

func envInitiator() {
	Envs = &EnvsSchema{
		HOST:                   viper.GetString("HOST"),
		PORT:                   viper.GetInt("PORT"),
		JWT_SECRET_KEY:         viper.GetString("JWT_SECRET_KEY"),
		JWT_EXP_HOURS:          viper.GetInt("JWT_EXP_HOURS"),
		JWT_REFRESH_SECRET_KEY: viper.GetString("JWT_REFRESH_SECRET_KEY"),
		JWT_REFRESH_EXP_HOURS:  viper.GetInt("JWT_REFRESH_EXP_HOURS"),
		MONGO_URI:              viper.GetString("MONGO_URI"),
		MONGO_DB_NAME:          viper.GetString("MONGO_DB_NAME"),
	}
}

func InitEnv(filepath string) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Warningf("error loading environment variables from %s: %w", filepath, err)
	}
	viper.AutomaticEnv()
	envInitiator()
}
