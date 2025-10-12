package config

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Unsetenv("PORT")
	port := GetPort()
	if port != "8080" {
		t.Errorf("expected default port 8080, got %s", port)
	}

	os.Setenv("PORT", "2222")
	defer os.Unsetenv("PORT")
	port = GetPort()
	if port != "2222" {
		t.Errorf("expected port 2222, got %s", port)
	}
}

func TestGetHostSigner(t *testing.T) {
	signer := GetHostSigner()
	if len(signer) == 0 {
		t.Error("expected host signer bytes to be non-empty")
	}
}
