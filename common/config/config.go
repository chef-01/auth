package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config structure holds application configurations.
type Config struct {
	AppEnv         string
	BaseURL        string
	MongoURI       string
	MongoUsername  string
	MongoPassword  string
	RedisAddr      string
	Port           string
	RequestTimeout int    // In seconds
	DatabaseName   string
	JwtSecret      string
	TokenDuration  int
	ApiVersion     string
	AllowedOrigins string
	LogPrefix      string
}

// LoadConfig loads environment variables from the appropriate .env file.
func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // Default to local if APP_ENV is not set.
	}

	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	timeout, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))
	if err != nil {
		timeout = 10 // default value if conversion fails
	}

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1"
	}

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*" // default to all origins
	}

	logPrefix := os.Getenv("LOG_PREFIX")
	if logPrefix == "" {
		logPrefix = "[PRIMUS] "
	}

	tokenDuration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	if err != nil {
		tokenDuration = 3600 // default value if conversion fails
	}

	return &Config{
		AppEnv:         os.Getenv("APP_ENV"),
		BaseURL:        os.Getenv("BASE_URL"),
		MongoURI:       os.Getenv("MONGO_URI"),
		MongoUsername:  os.Getenv("MONGO_USERNAME"),
		MongoPassword:  os.Getenv("MONGO_PASSWORD"),
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		Port:           os.Getenv("PORT"),
		RequestTimeout: timeout,
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		TokenDuration:  tokenDuration,
		ApiVersion:     apiVersion,
		AllowedOrigins: allowedOrigins,
		LogPrefix:      logPrefix,
	}
}