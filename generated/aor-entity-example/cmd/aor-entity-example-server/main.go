package main

import (
	// "flag"
	"fmt"
	"log"

	"github.com/roscopecoltran/aor-gin-swagger/config"
	"github.com/roscopecoltran/aor-gin-swagger/database"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi"
)

var (
	apiConfig restapi.Config
	// configFile string
)

func main() {

	// configFile := flag.String("conf", "config.yml", "configuration file")

	if config.Config.Debug {
		fmt.Printf("Listening on: %v\n", config.Config.Port)
		fmt.Printf("Database status: %t\n", database.Status)
	}

	err := apiConfig.Parse()
	if err != nil {
		log.Fatal(err)
	}

	database.CreateTables(true, swaggerGeneratedModels...)
	svc := &AorEntityExampleServer{Health: false}
	api := restapi.NewAPI(svc, &apiConfig)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}

}
