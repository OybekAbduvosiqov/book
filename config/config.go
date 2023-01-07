package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {

	cfg := Config{}

	cfg.HTTPPort = ":8800"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "oybek"
	cfg.PostgresDatabase = "catalog"
	cfg.PostgresPassword = "oybek"
	cfg.PostgresPort = "5432"

	return cfg
}
