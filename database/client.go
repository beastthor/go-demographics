package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

// connector variable used for CRUD operations
var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {

	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection was succesful!!")
	return nil

}
