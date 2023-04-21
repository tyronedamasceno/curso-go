package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))
}

func RetrieveUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Retrieving Users"))
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Retrieving User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
