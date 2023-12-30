package utils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

const minSecretLength = 32

var envVars = struct {
	ApiPort      string
	ApiJwtSecret string
}{ApiPort: "API_PORT", ApiJwtSecret: "API_JWT_SECRET"}

type AppConfig struct {
	Port      int
	JwtSecret []byte
}

var appConfig *AppConfig

func (config *AppConfig) setApiPort(errSlice *[]string) {
	port, portErr := strconv.Atoi(os.Getenv(envVars.ApiPort))

	if nil != portErr {
		*errSlice = append(*errSlice, fmt.Sprintf("- %s must be a valid port", envVars.ApiPort))
		return
	}

	config.Port = port
}

func (config *AppConfig) setApiJwtSecret(errSlice *[]string) {
	secret := os.Getenv(envVars.ApiJwtSecret)

	if "" == secret {
		*errSlice = append(*errSlice, fmt.Sprintf("- %s cannot be empty", envVars.ApiJwtSecret))
		return
	}

	if minSecretLength > len(secret) {
		*errSlice = append(*errSlice, fmt.Sprintf("- %s must be at least %d characters", envVars.ApiJwtSecret, minSecretLength))
		return
	}

	config.JwtSecret = []byte(secret)
}

func parseEnv() (*AppConfig, error) {
	dotenvErr := godotenv.Load()
	var parseErr []string

	if nil != dotenvErr {
		panic(dotenvErr)
	}

	config := AppConfig{}

	config.setApiPort(&parseErr)
	config.setApiJwtSecret(&parseErr)

	if nil == parseErr {
		return &config, nil
	}

	return nil, errors.New(fmt.Sprintf("\n%s", strings.Join(parseErr[:], "\n")))
}

func GetAppConfig() *AppConfig {
	if nil != appConfig {
		return appConfig
	}

	config, err := parseEnv()

	if nil != err {
		log.Fatal(err)
	}

	appConfig = config
	return appConfig
}
