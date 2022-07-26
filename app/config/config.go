package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	Debug bool `envconfig:"DEBUG"`

	//! Server
	Server struct {
		Port    string `envconfig:"SERVER_PORT"`
		Timeout int    `envconfig:"SERVER_TIMEOUT"`
	}

	//! MYSQL
	Mysql struct {
		Host string `envconfig:"MYSQL_DB_HOST"`
		Port string `envconfig:"MYSQL_DB_PORT"`
		User string `envconfig:"MYSQL_DB_USER"`
		Pass string `envconfig:"MYSQL_DB_PASS"`
		Name string `envconfig:"MYSQL_DB_NAME"`
	}

	//! JWT
	JWT struct {
		Secret  string `envconfig:"JWT_SECRET"`
		Expired int    `envconfig:"JWT_EXPIRED"`
	}
}

func GetConfig() Config {
	var conf Config

	var (
		filename = os.Getenv("CONFIG_FILE")
	)

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &conf); err != nil {
			panic(err)
		}
		panic(err)
	}

	if err := godotenv.Load(filename); err != nil {
		panic(err)
	}

	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}
	return conf
}
