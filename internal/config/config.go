package config

import (
	"log"
	"os"
	"sync"
	"todo/internal/config/project"
	"todo/internal/config/server"

	"github.com/ilyakaznacheev/cleanenv"
)

func getConfigPath() string { return "./config/app/main.yaml" }

func configPathValidate(configPath *string) {
	if *configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", *configPath)
	}
}

func setConfig(config *Config, configPath *string) {
	if err := cleanenv.ReadConfig(*configPath, config); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}
}

type Config struct {
	Project project.Config `yaml:"project"`
	Server  server.Config  `yaml:"server"`
}

func GetConfig() *Config {
	var config Config
	var once sync.Once

	once.Do(func() {
		configPath := getConfigPath()
		configPathValidate(&configPath)
		setConfig(&config, &configPath)
	})

	return &config
}
