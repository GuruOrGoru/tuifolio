package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/guruorgoru/tuifolio/internal/config"
	handlers "github.com/guruorgoru/tuifolio/internal/ssh"
)

func main() {
	port, err := config.GetPort()
	if err != nil {
		log.Error(err)
	}
	host, err := config.GetHost()
	if err != nil {
		log.Error(err)
	}

	signer := config.GetHostSigner()
	server, err := handlers.NewSSHServer(port, host, signer)
	if err != nil {
		log.Error(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Info("Starting SSH server", "host", host, "port", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error(err)
		}
	}()

	<-done
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Info("Server exited properly")
}
