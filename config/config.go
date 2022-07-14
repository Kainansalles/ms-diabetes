package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// InitEnvConfig inicia file de env.
func InitEnvConfig() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Error loading Env: %s \n", envErr.Error())
	}
}
