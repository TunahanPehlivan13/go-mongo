package main

import (
	"github.com/TunahanPehlivan13/go-mongo/config"
	"github.com/TunahanPehlivan13/go-mongo/server"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	app := server.NewApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("app.port")
	}
	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
