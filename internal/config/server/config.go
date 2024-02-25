package server

import (
	"time"
)

type Config struct {
	Port        string        `yaml:"port" env:"PORT" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-required:"true"`
}
