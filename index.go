package main

import (
	"fmt"
	"github.com/omar-ozgur/gram/config"
	"github.com/omar-ozgur/gram/db"
	"github.com/omar-ozgur/gram/utilities"
	"net/http"
)

func main() {
	args := config.ParseArgs()

	db.InitDB(args.Service)

	n := config.InitRouter()

	utilities.Sugar.Infof("Started server on port %s\n", args.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", args.Port), n)
	if err != nil {
		utilities.Logger.Fatal(err.Error())
	}
}
