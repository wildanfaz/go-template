package configs

import (
	"os"
	"time"
)

type Config struct {
	EchoPort    string
	JWTSecret   []byte
	JWTDuration time.Duration
}

func InitConfig() *Config {
	jwtDuration, err := time.ParseDuration(GetEnv("JWT_DURATION", "1h"))
	if err != nil {
		panic(err)
	}

	return &Config{
		EchoPort:    GetEnv("ECHO_PORT", ":1323"),
		JWTSecret:   []byte(GetEnv("JWT_SECRET", "secret")),
		JWTDuration: jwtDuration,
	}
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
