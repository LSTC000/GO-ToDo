package mongo

type Config struct {
	Host     string `yaml:"host" env:"MONGO_HOST" env-required:"true"`
	Port     string `yaml:"port" env:"MONGO_PORT" env-required:"true"`
	Database string `yaml:"database" env:"MONGO_INITDB_DATABASE" env-required:"true"`
	Username string `yaml:"username" env:"MONGO_INITDB_ROOT_USERNAME" env-required:"true"`
	Password string `yaml:"password" env:"MONGO_INITDB_ROOT_PASSWORD" env-required:"true"`
}
