package main

import (
	"log"

	"github.com/guruorgoru/tuifolio/internal/config"
	handlers "github.com/guruorgoru/tuifolio/internal/ssh"
)

func main() {
	port := config.GetPort()
	signer := config.GetHostSigner()
	server, err := handlers.NewSSHServer(port, signer)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Server listening at port %v\n", port)
	log.Fatalln(server.ListenAndServe())
}
