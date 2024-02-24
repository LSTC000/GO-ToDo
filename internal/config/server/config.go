package server

import (
	"time"
)

type Config struct {
	Host        string        `yaml:"host" env:"HOST" env-required:"true"`
	Port        string        `yaml:"port" env:"PORT" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-required:"true"`
}
