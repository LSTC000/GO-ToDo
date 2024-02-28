package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"todo/internal/config/project"
	"todo/internal/config/server"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Project project.Config `yaml:"project"`
	Server  server.Config  `yaml:"server"`
	Mongo   server.Config  `yaml:"mongo"`
}

func getConfigPath() string { return "./config/app/main.yaml" }

func configPathValidate(configPath *string) error {
	if *configPath == "" {
		return fmt.Errorf("not config by path: %s", *configPath)
	}
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

func setConfig(config *Config, configPath *string) error {
	if err := cleanenv.ReadConfig(*configPath, config); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	var config Config
	var once sync.Once

	once.Do(func() {
		configPath := getConfigPath()

		if err := configPathValidate(&configPath); err != nil {
			log.Fatalf("Config path error: %v", err)
		}

		if err := setConfig(&config, &configPath); err != nil {
			log.Fatalf("Config create error: %v", err)
		}
	})

	return &config
}
