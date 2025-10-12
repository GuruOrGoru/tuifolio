package config

import (
	_ "embed"
	"os"

	"github.com/joho/godotenv"
)

func GetPort() string {
	_ = godotenv.Load(".env") // Ignore error if .env not found
	portstr := os.Getenv("PORT")
	if portstr == "" {
		portstr = "8080"
	}
	return portstr
}

//go:embed host_key
var hostKeyBytes []byte

func GetHostSigner() []byte {
	return hostKeyBytes
}
