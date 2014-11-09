package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	. "goboilerplate/app/models"
	"goboilerplate/config"
)

type DatabaseConnection struct {
	*gorm.DB
}

func NewDatabaseConnection() (*DatabaseConnection, error) {
	db, err := gorm.Open("postgres", config.PostgresArgs)
	if err != nil {
		return nil, err
	}
	return &DatabaseConnection{&db}, nil
}

func Migrate() {
	db, err := NewDatabaseConnection()
	if err != nil {
		fmt.Printf("Error migrating database: %s\n", err.Error())
		return
	}

	if err := db.AutoMigrate(&Cat{}).Error; err != nil {
		fmt.Printf("Error migrating database: %s\n", err.Error())
	}
}

func main() {
	Migrate()
}
