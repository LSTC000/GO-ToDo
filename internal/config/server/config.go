package server

type Config struct {
	Port string `yaml:"port" env:"SERVER_PORT" env-required:"true"`
}
