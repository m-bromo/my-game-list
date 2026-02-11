package main

import (
	"log"

	"github.com/m-bromo/my-game-list/config"
	"github.com/m-bromo/my-game-list/pkg/logging"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log := logging.NewLogger(config)

	log.Log.Info("starting application")

}
