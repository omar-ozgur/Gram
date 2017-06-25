package config

import (
	"flag"
	"github.com/omar-ozgur/gram/utilities"
	"os"
)

type Args struct {
	Port    string
	Service string
}

func GetPort() string {
	port := os.Getenv("GRAM_PORT")

	// Set default port if there is nothing in the environment
	if port == "" {
		port = utilities.DefaultPort
	}

	return port
}

func ParseArgs() *Args {
	var args Args

	flag.StringVar(&args.Port, "port", GetPort(), "Specifies the port for the server to run on. Ex: --port 3000")
	flag.StringVar(&args.Service, "service", "default", "Specifies the name of the service that the authentication server is being used for. Gram supports users for multiple services. Ex: --service MY_SERVICE")

	flag.Parse()

	return &args
}
