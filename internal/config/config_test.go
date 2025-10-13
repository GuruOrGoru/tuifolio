package config

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Unsetenv("PORT")
	_, err := GetPort()
	if err == nil {
		t.Error("expected error when PORT not set")
	}

	os.Setenv("PORT", "2222")
	defer os.Unsetenv("PORT")
	port, err := GetPort()
	if err != nil || port != "2222" {
		t.Errorf("expected port 2222, got %s", port)
	}
}

func TestGetHostSigner(t *testing.T) {
	signer := GetHostSigner()
	if len(signer) == 0 {
		t.Error("expected host signer bytes to be non-empty")
	}
}
