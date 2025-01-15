package main

import (
	"github.com/alhaos/invertMultiplicationTrainer/internal/appConfig"
	"github.com/alhaos/invertMultiplicationTrainer/internal/webServer"
)

func main() {

	// Init appConfig
	config, err := appConfig.New()
	if err != nil {
		panic(err)
	}

	// Init web server
	ws := webServer.New(config.WebServer)

	// Run web server
	err = ws.Run()
	if err != nil {
		panic(err)
	}

}
