// Package config contains functionality to load environment variables and other configuration.
package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	keyEnvironment = "ENVIRONMENT"
	keyServiceName = "SERVICE_NAME"
)

// Configuration contains the configuration for the service.
type Configuration struct {
	Environment string
	ServiceName string
}

// LoadConfiguration is used to load the configuration from the environment variables provided.
func LoadConfiguration() Configuration {
	config := Configuration{
		Environment: getOrDefault(keyEnvironment, "local", ""),
		ServiceName: getOrDefault(keyServiceName, "gogrumpy-game-provider", ""),
	}

	return config
}

func getOrDefault(key string, defaultValue string, defaultCommand string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		if len(defaultCommand) != 0 {
			startCmd(defaultCommand)
			time.Sleep(500 * time.Millisecond) // a port forward may take some time to complete
		}

		return defaultValue
	}

	return val
}

func startCmd(cmd string) {
	fmt.Println("Execute " + cmd)
	cmdsh := exec.Command("sh", "-c", cmd)

	var err error
	if strings.Contains(cmd, "port-forward") {
		err = cmdsh.Start()
	} else {
		err = cmdsh.Run()
	}

	if err != nil {
		log.Println("Could not execute " + cmd + ": " + err.Error())
	}
}
