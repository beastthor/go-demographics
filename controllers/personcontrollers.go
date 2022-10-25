package controllers

import (
	"encoding/json"
	"go-demographics/database"
	"go-demographics/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllPerson get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var person []entity.Person
	database.Connector.Find(&person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)

}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)

}
func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
