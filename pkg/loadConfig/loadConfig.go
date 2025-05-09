package load_config

import (
	"gin/internal/config"
	"gin/pkg/database"
	"gin/pkg/get_env"
)

func LoadConfig() config.Config {
	return config.Config{
		BindAddr: get_env.GetEnv("BIND_ADDR", "8080"),
		Connection: database.Config{
			Host:     get_env.GetEnv("DB_HOST", "localhost"),
			Port:     5432,
			User:     get_env.GetEnv("DB_USER", "myuser"),
			Password: get_env.GetEnv("DB_PASSWORD", "mypassword"),
			DBName:   get_env.GetEnv("DB_NAME", "mydb"),
			SSLMode:  get_env.GetEnv("DB_SSLMODE", "disable"),
			TimeZone: get_env.GetEnv("DB_TIMEZONE", "UTC"),
		},
	}
}
