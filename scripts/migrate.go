package main

import (
	"fmt"
	_ "goboilerplate/app/models"
	. "goboilerplate/db"
)

func main() {
	fmt.Println("Beginning migration")
	numMigrated, err := Migrate()
	if err != nil {
		fmt.Println("Error migrating database")
	} else {
		fmt.Printf("Migrated %d models\n", numMigrated)
	}
}
