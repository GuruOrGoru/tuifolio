package config

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	err := os.Unsetenv("PORT")
	if err != nil {
		t.Fatalf("failed to unset PORT: %v", err)
	}
	_, err = GetPort()
	if err == nil {
		t.Error("expected error when PORT not set")
	}

	err = os.Setenv("PORT", "2222")
	if err != nil {
		t.Fatalf("failed to set PORT: %v", err)
	}
	defer func() {
		err := os.Unsetenv("PORT")
		if err != nil {
			t.Fatalf("failed to unset PORT: %v", err)
		}
	}()
	port, err := GetPort()
	if err != nil || port != "2222" {
		t.Errorf("expected port 2222, got %s", port)
	}
}

func TestGetHostSigner(t *testing.T) {
	signer := GetHostSigner()
	if len(signer) == 0 {
		t.Error("expected non-empty host signer")
	}
}
