package project

type Config struct {
	Mode string `yaml:"mode" env:"MODE" env-required:"true" env-description:"local/dev/prod"`
}
