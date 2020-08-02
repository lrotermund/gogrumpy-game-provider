package main

import (
	"fmt"

	"github.com/lrotermund/gogrumpy-game-provider/pkg/api"
	"github.com/lrotermund/gogrumpy-game-provider/pkg/config"
)

const (
	address = ":8080"
)

func main() {
	config := config.LoadConfiguration()

	fmt.Printf("Starting %s in %s mode\n", config.ServiceName, config.Environment)

	e, err := api.New(config)
	if err != nil {
		panic(err)
	}

	e.Logger.Error(e.Start(address))
}
