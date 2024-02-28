package project

type Config struct {
	Mode string `yaml:"mode" env:"PROJECT_MODE" env-required:"true" env-description:"local/dev/prod"`
}
