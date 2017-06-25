package main

import (
	"fmt"
	"github.com/omar-ozgur/gram/app/models"
	"github.com/omar-ozgur/gram/config"
	"github.com/omar-ozgur/gram/db"
	"github.com/omar-ozgur/gram/utilities"
	"net/http"
)

func main() {
	config.ParseArgs()

	db.InitDB()

	models.Init()

	n := config.InitRouter()

	utilities.Sugar.Infof("Started server on port %s\n", utilities.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", utilities.Port), n)
	if err != nil {
		utilities.Logger.Fatal(err.Error())
	}
}
