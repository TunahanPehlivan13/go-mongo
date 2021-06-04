package main

import (
	"github.com/TunahanPehlivan13/go-mongo/server"
	"log"
	"os"
)

func main() {
	app := server.NewApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
