// /users - GET
// /users/{id} - GET
// /users - POST
// /users/{id} - PUT
// /users/{id} - DELETE

package main

import (
	"log"
	"net/http"

	"github.com/Roman77St/simple_project/storage"
	"github.com/Roman77St/simple_project/rest_api/handlers"

	"github.com/gorilla/mux"
)


func main()  {
	err := storage.InitDatabase("./storage/sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}