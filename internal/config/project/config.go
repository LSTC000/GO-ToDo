package project

type Config struct {
	Version string `yaml:"version" env:"VERSION" env-required:"true"`
	Name    string `yaml:"name" env:"NAME" env-required:"true"`
	Mode    string `yaml:"mode" env:"MODE" env-required:"true" env-description:"local/dev/prod"`
}
