package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost   string
	AppEnv       string
	AppDebug     string
	ApiPort      string
	DBUsername   string
	DBPort       string
	DBPassword   string
	DBHost       string
	DBName       string
	JWTSecret    string
	JWTExpire    int
	BycrptSalt   string
	S3Region     string
	S3Id         string
	S3SecretKey  string
	S3BucketName string
	AppUrl       string
}

// NOTE: Global Config Variable
var ConfigEnv = InitConfig()

func InitConfig() Config {

	if err := godotenv.Load(); err != nil {
		// Just log, don't panic
		fmt.Println("[WARNING] .env file not found or could not be loaded. Proceeding with system environment variables.")
	}

	return Config{
		PublicHost:   GetEnv("PUBLIC_HOST", "localhost"),
		AppEnv:       GetEnv("APP_ENV", "development"),
		AppDebug:     GetEnv("APP_DEBUG", "true"),
		ApiPort:      GetEnv("API_PORT", "8090"),
		DBUsername:   GetEnv("DB_USERNAME", "root"),
		DBPort:       GetEnv("DB_PORT", "5432"),
		DBPassword:   GetEnv("DB_PASSWORD", ""),
		DBHost:       GetEnv("DB_HOST", "localhost"),
		DBName:       GetEnv("DB_NAME", ""),
		JWTSecret:    GetEnv("JWT_SECRET", "secret"),
		JWTExpire:    GetEnvInteger("JWT_EXPIRE", 3600*2), // Default to 2 hours
		BycrptSalt:   GetEnv("BCRYPT_SALT", "10"),
		S3Region:     GetEnv("S3_REGION", ""),
		S3Id:         GetEnv("S3_ID", ""),
		S3SecretKey:  GetEnv("S3_SECRET_KEY", ""),
		S3BucketName: GetEnv("S3_BUCKET_NAME", ""),
		AppUrl:       GetEnv("APP_URL", "http://localhost:8089"),
	}
}

func GetEnv(key string, defaultValue string) string {
	val := defaultValue
	if value, ok := os.LookupEnv(key); ok {
		val = value
	}
	return val
}

func GetEnvInteger(key string, defaultValue int) int {
	val := defaultValue
	if value, ok := os.LookupEnv(key); ok {
		valInt, _ := strconv.ParseInt(value, 10, 64)

		return int(valInt)
	}

	return val
}
