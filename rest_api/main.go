package main

import (
	"log"
	"net/http"

	"github.com/Roman77St/simple_project/rest_api/handlers"
	"github.com/Roman77St/simple_project/storage"

	"github.com/gorilla/mux"
)


func main()  {
	err := storage.InitDatabase()
	if err != nil {
		log.Fatal("Ошибка подключения к реляционной базе данных: ", err)
	}
	err = storage.InitNewClient("redis:6379") // "redis:6379" - compose, "localhost:6379" - go run
	if err != nil {
		log.Fatal("Ошибка подключения к Redis: ", err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	http.ListenAndServe(":8080", router)
}