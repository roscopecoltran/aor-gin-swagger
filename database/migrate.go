package database

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"
	"github.com/k0kubun/pp"
	"github.com/qianlnk/exception"
	"github.com/roscopecoltran/aor-gin-swagger/config"
)

func CreateTables(isTruncate bool, models ...interface{}) {
	config.PrintConfig()
	printDatabaseInfo()
	fmt.Println("models count: ", len(models))
	for _, model := range models {
		isStruct := structs.IsStruct(model)
		modelType := getType(model)
		if isTruncate && isStruct {
			if err := DB.DropTableIfExists(model).Error; err != nil {
				/*
					if config.Config.StrictMode {
						panic(err)
					}
				*/
				fmt.Printf("error occured while dropping table: %s / %s \n", structs.Name(model), modelType)
			}
		}
		if isStruct {
			fmt.Printf("migration model: %s / %s \n", structs.Name(model), modelType)
			tr := exception.New()
			tr.Try(
				func() {
					DB.AutoMigrate(model)
				},
			).Catch(
				func(e exception.Exception) {
					fmt.Println("exception:", e)
				},
			)
		} else {
			fmt.Printf("skip Migration model '%s' as not a valid.\n", modelType)
		}
	}
}

func getType(object interface{}) string {
	if t := reflect.TypeOf(object); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func printDatabaseInfo() {
	dbInfo := structs.Map(config.Config.DB)
	pp.Println("Database Info: ", dbInfo)

}
