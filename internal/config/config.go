package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type HTTPAddress struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Database struct {
	Host     string `yamal:"host" env-required:"true"`
	Port     uint	`yaml:"port" env-required:"true"`
	Name     string	`yaml:"name" env-required:"true"`
	Username string	`yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	SSLMode  string `yaml:"sslmode" envrequired:"true"`
}

type Config struct {
	Env string `yaml:"env" env-required:"true"`
	HTTPAddress `yaml:"http_server" env-required:"true"`
	Database	`yaml:"database" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("lenv file not found")
	}

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "config file path")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("config file path not provided")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Error loading config file: %v", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return &cfg
}
