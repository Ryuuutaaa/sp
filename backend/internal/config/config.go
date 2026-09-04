package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	GQLURL          string
	R2AccountID     string
	R2AccessKeyID   string
	R2SecretKey     string
	R2BucketName    string
	R2PublicURL     string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	return &Config{
		Port:          getEnv("PORT", "8080"),
		GQLURL:        getEnv("GQL_URL", "http://localhost:4000/graphql"),
		R2AccountID:   getEnv("R2_ACCOUNT_ID", ""),
		R2AccessKeyID: getEnv("R2_ACCESS_KEY_ID", ""),
		R2SecretKey:   getEnv("R2_SECRET_ACCESS_KEY", ""),
		R2BucketName:  getEnv("R2_BUCKET_NAME", ""),
		R2PublicURL:   getEnv("R2_PUBLIC_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
