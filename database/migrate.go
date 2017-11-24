package database

import (
	"fmt"

	"github.com/fatih/structs"
)

func CreateTables(isTruncate bool, tables ...interface{}) {
	for _, table := range tables {
		if isTruncate {
			if err := DB.DropTableIfExists(table).Error; err != nil {
				panic(err)
			}
		}
		fmt.Printf("AutoMigrate, table: %s\n", structs.Name(table))
		DB.AutoMigrate(table)
	}
}
