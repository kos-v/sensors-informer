package main

import (
	"flag"
	informerApp "github.com/kos-v/sensors-informer/internal/app"
	"log"
)

func main() {
	specificConfig := flag.String("config", "", "The path to a specific configuration file")
	flag.Parse()

	app := informerApp.NewApp(*specificConfig)
	if err := app.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
