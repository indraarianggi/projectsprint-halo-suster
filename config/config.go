package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost      string `mapstructure:"APP_HOST"`
	AppPort      string `mapstructure:"APP_PORT"`
	DBName       string `mapstructure:"DB_NAME"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBUsername   string `mapstructure:"DB_USERNAME"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBParams     string `mapstructure:"DB_PARAMS"`
	DBScehma     string `mapstructure:"DB_SCHEMA"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	BcryptSalt   string `mapstructure:"BCRYPT_SALT"`
	S3AccessKey  string `mapstructure:"AWS_ACCESS_KEY_ID"`
	S3SecretKey  string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	S3BucketName string `mapstructure:"AWS_S3_BUCKET_NAME"`
	S3Region     string `mapstructure:"AWS_REGION"`
}

func Load() (config Config) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, error:", err.Error())
	}

	config = Config{
		AppHost:      os.Getenv("APP_HOST"),
		AppPort:      os.Getenv("APP_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		DBPort:       os.Getenv("DB_PORT"),
		DBHost:       os.Getenv("DB_HOST"),
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBParams:     os.Getenv("DB_PARAMS"),
		DBScehma:     os.Getenv("DB_SCHEMA"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		BcryptSalt:   os.Getenv("BCRYPT_SALT"),
		S3AccessKey:  os.Getenv("AWS_ACCESS_KEY_ID"),
		S3SecretKey:  os.Getenv("AWS_SECRET_ACCESS_KEY"),
		S3BucketName: os.Getenv("AWS_S3_BUCKET_NAME"),
		S3Region:     os.Getenv("AWS_REGION"),
	}

	return
}

func (config *Config) GetDSN() (dsn string) {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s&search_path=%s",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBParams,
		config.DBScehma)
}
