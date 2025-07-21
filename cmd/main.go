package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse-app ", log.LstdFlags)
	serv := server.Serv(logger)
	err := serv.Serv.ListenAndServe()
	if err != nil {
		logger.Fatalf("mistake with server: %v", err)
	}

}
