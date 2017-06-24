package config

import (
	"flag"
	"os"
)

type Args struct {
	Port string
}

func GetPort() string {
	defaultPort := "3000"

	port := os.Getenv("GRAM_PORT")

	// Set default port if there is nothing in the environment
	if port == "" {
		port = defaultPort
	}

	return port
}

func ParseArgs() Args {
	var args Args

	flag.StringVar(&args.Port, "port", GetPort(), "Specifies the port for the server to run on. Ex: --port 3000")

	flag.Parse()

	return args
}
