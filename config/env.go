package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT						string `mapstructure:"PORT"`
	URL 						string `mapstructure:"URL"`
	MONGODB_CONNECTION_URL		string `mapstructure:"MONGODB_CONNECTION_URL"`
	MONGODB_DATABASE 			string `mapstructure:"MONGODB_DATABASE"`
	AUTH_SECRET					string `mapstructure:"AUTH_SECRET"`
	AUTH_DATABASE				string `mapstructure:"AUTH_DATABASE"`
	TOKEN_EXPIRY				int    `mapstructure:"TOKEN_EXPIRY"`
}

func LoadEnv() (*Env){
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Env file not loaded. Here's what happened : %v ", err)
		panic(err)
	}

	TokenExpiry, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRY"))

	if err != nil {
		log.Printf("Env file not loaded. Here's what happened : %v ", err)
		panic(err)
	}

	return &Env{
		PORT: os.Getenv("PORT"),
		URL: os.Getenv("URL"),
		MONGODB_CONNECTION_URL: os.Getenv("MONGODB_CONNECTION_URL"),
		MONGODB_DATABASE: os.Getenv("MONGODB_DATABASE"),
		AUTH_SECRET: os.Getenv("AUTH_SECRET"),
		AUTH_DATABASE: os.Getenv("AUTH_DATABASE"),
		TOKEN_EXPIRY: TokenExpiry,
	}
}