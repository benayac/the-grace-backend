package pkg

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

type Config struct {
	DbHost            string
	DbPort            int
	DbUsername        string
	DbPassword        string
	DbName            string
	Email             string
	EmailPassword     string
	RedisHost         string
	SigningKey        string
	SigningKeyEncrypt string
}

var Conf *Config

func GetConfigEnv() error {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}
	Conf = &Config{
		DbHost:            os.Getenv("DB_HOST"),
		DbPort:            port,
		DbUsername:        os.Getenv("DB_USERNAME"),
		DbPassword:        os.Getenv("DB_PASSWORD"),
		DbName:            os.Getenv("DB"),
		Email:             os.Getenv("EMAIL"),
		EmailPassword:     os.Getenv("EMAIL_PASSWORD"),
		RedisHost:         os.Getenv("REDIS_HOST"),
		SigningKey:        os.Getenv("SIGNING_KEY"),
		SigningKeyEncrypt: os.Getenv("SIGNING_KEY_ENCRYPT"),
	}
	return nil
}

func GetConfigJson() error {
	viper.SetConfigType("json")
	viper.AddConfigPath("../../")
	viper.SetConfigName("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	Conf = &Config{
		DbHost:            viper.GetString("dbHost"),
		DbPort:            viper.GetInt("dbPort"),
		DbUsername:        viper.GetString("dbUsername"),
		DbPassword:        viper.GetString("dbPassword"),
		DbName:            viper.GetString("dbName"),
		Email:             viper.GetString("email"),
		EmailPassword:     viper.GetString("emailPassword"),
		RedisHost:         viper.GetString("redisHost"),
		SigningKey:        viper.GetString("signingKey"),
		SigningKeyEncrypt: os.Getenv("SIGNING_KEY_ENCRYPT"),
	}
	return nil
}
