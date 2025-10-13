package config

import (
	_ "embed"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetPort() (string, error) {
	_ = godotenv.Load(".env")
	portstr := os.Getenv("PORT")
	if portstr == "" {
		return "", errors.New("ENV Port not set in .env")
	}
	return portstr, nil
}

func GetHost() (string, error) {
	host := os.Getenv("HOST")
	if host == "" {
		return "", errors.New("ENV Host not set in .env")
	}
	return host, nil
}

//go:embed host_key
var hostKeyBytes []byte

func GetHostSigner() []byte {
	return hostKeyBytes
}
