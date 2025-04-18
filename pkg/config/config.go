package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST             string
	DB_PORT             string
	DB_USER             string
	DB_NAME             string
	DB_PASSWORD         string
	EDU_TEST            string
	MINIO_ROOT_USER     string
	MINIO_ROOT_PASSWORD string
	MINIO_HOST          string
	MINIO_PUBLIC_HOST   string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("error loading .env file or not found", err)
	}

	config := Config{}

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "education_center"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "secret"))
	config.EDU_TEST = cast.ToString(coalesce("EDU_TEST", "8080"))
	config.MINIO_ROOT_USER = cast.ToString(coalesce("MINIO_ROOT_USER", ""))
	config.MINIO_ROOT_PASSWORD = cast.ToString(coalesce("MINIO_ROOT_PASSWORD", ""))
	config.MINIO_HOST = cast.ToString(coalesce("MINIO_HOST", "minio:9000"))
	config.MINIO_PUBLIC_HOST = cast.ToString(coalesce("MINIO_PUBLIC_HOST", ""))
	return config
}

func coalesce(key string, defValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defValue
}
