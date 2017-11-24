package main

import (
	"fmt"
	"log"

	"github.com/roscopecoltran/gin-swagger/config"
	"github.com/roscopecoltran/gin-swagger/database"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi"
)

var (
	apiConfig restapi.Config
)

func main() {

	if config.Config.Debug {
		fmt.Printf("Listening on: %v\n", config.Config.Port)
		fmt.Printf("Database status: %t\n", database.Status)
	}

	err := apiConfig.Parse()
	if err != nil {
		log.Fatal(err)
	}

	database.CreateTables(true, database.DefaultTables...)
	svc := &AorEntityServer{Health: false}

	api := restapi.NewAPI(svc, &apiConfig)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}

}
