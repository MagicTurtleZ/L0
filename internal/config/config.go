package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address 	string			`yaml:"address" env-default:"localhost:8080"`
	Timeout 	time.Duration	`yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration	`yaml:"idle-timeout" env-defaut:"45s"`
}

type NATSettings struct {
	ClusterId string	`yaml:"cluster_id" env-default:"world-nats-stage"`
	ClientId  string	`yaml:"client_id" env-default:"woonbeaj"`
}

type Config struct {
	Env        string	`yaml:"env" env-default:"local"`
	StorageURL string	`yaml:"storage_url" env-required:"true"`
	HTTPServer		 	`yaml:"http_server"`
	NATSettings		 	`yaml:"nuts_settings"`
}

func MustLoad(cfgPath string) *Config {
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", cfgPath)
	}
	var cfg Config 
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("canno`t read config file: %s", cfgPath)
	}

	return &cfg
}