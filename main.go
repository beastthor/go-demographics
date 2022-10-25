package main

import (
	"go-demographics/controllers"
	"go-demographics/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	initDB()

	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))

}

type Person struct {
	FirstName string
	LastName  string
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

func initDB() {
	config := database.Config{
		ServerName: "127.0.0.1:3306",
		User:       "root",
		Password:   "password",
		DB:         "aadharcard",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}
}
