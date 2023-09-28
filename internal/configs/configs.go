package configs

import (
	"os"
	"speech-model-hub/internal/utils"

	"github.com/spf13/viper"
)

type AppConfig struct {
	MongoDB     MongoConfig `name:"mongodb" mapstructure:"mongodb" validate:"required"`
	Port        string      `name:"port" mapstructure:"port" validate:"required"`
	environment string
}

type MongoConfig struct {
	MongoDBURL   string `name:"mongo_url" mapstructure:"mongo_url" validate:"required"`
	DatabaseName string `name:"database_name" mapstructure:"database_name" validate:"required"`
}

func loadConfig(path string) AppConfig {
	env := os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		viper.SetConfigName("env.production")
	} else if env == "STAGING" {
		viper.SetConfigName("env.staging")
	} else {
		env = "LOCAL"
		viper.SetConfigName("env.local")
	}
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	config.environment = env
	return config
}

func (config *AppConfig) Validate() error {
	return utils.ValidateStruct(config)
}

func (config *AppConfig) IsProduction() bool {
	return config.environment == "PRODUCTION"
}

func (config *AppConfig) IsStaging() bool {
	return config.environment == "STAGING"
}

func (config *AppConfig) IsLocal() bool {
	return config.environment == "LOCAL"
}

func NewConfig() AppConfig {
	appConfig := loadConfig("./")
	err := appConfig.Validate()
	if err != nil {
		panic(err)
	}
	return appConfig
}

func NewTestConfig(path string) AppConfig {
	appConfig := loadConfig(path)
	err := appConfig.Validate()
	if err != nil {
		panic(err)
	}
	return appConfig
}
