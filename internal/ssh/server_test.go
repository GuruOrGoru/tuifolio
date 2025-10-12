package handlers

import (
	"testing"

	"github.com/guruorgoru/tuifolio/internal/config"
)

func TestNewSSHServer(t *testing.T) {
	port := "2222"
	signer := config.GetHostSigner()

	server, err := NewSSHServer(port, signer)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if server == nil {
		t.Error("expected server to be non-nil")
	}
}
