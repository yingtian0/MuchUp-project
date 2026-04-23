package config

import "os"

type Config struct {
	SecretKey string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	GRPCPort  string
	HTTPPort  string
	WSPort    string
	RedisAddr string
}

func LoadConfig() *Config {
	return &Config{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
		DBHost:    os.Getenv("DB_HOST"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
		DBPort:    os.Getenv("DB_PORT"),
		GRPCPort:  getEnv("GRPC_PORT", "9000"),
		HTTPPort:  getEnv("SERVER_PORT", "8080"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func (c *Config) GetRedisAddr() string {
	if c.RedisAddr == "" {
		return "localhost:6379"
	}
	return c.RedisAddr
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
