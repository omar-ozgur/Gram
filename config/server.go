package config

import (
	"github.com/omar-ozgur/gram/utilities"
	"os"
)

var defaultPort = "3000"

type Args struct {
	Port string
}

func GetPort() string {
	var port = os.Getenv("PORT")

	// Set a default port if there is nothing in the environment
	if port == "" {
		port = defaultPort
		utilities.Sugar.Infof("No PORT environment variable detected, defaulting to port %s", port)
	}
	return port
}

func ParseArgs() Args {
	var args Args

	// TODO: Attempt to get port from command line arguments
	// Else:
	args.Port = GetPort()

	// TODO: Parse other arguments

	return args
}
