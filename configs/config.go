package configs

import "os"

type Config struct {
	EchoPort string
}

func InitConfig() *Config {
	echoPort := GetEnv("ECHO_PORT", ":1323")

	return &Config{
		EchoPort: echoPort,
	}
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
