package config

import "os"

var (
	ListenPort = getEnv("LISTEN_PORT", "3061")
	MysqlDsn   = getEnv("MYSQL_DSN", "root:root123456@(localhost:3306)/FilmSite?charset=utf8mb4&parseTime=True&loc=Local")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
