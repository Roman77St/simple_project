package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Roman77St/simple_project/rest_api/handlers"
	"github.com/Roman77St/simple_project/storage"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users", handlers.CreateUser).Methods("POST", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:8001"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешенные методы
		AllowedHeaders: []string{"Content-Type", "Authorization"}, // Разрешенные заголовки
		ExposedHeaders: []string{"Content-Length"}, // Заголовки, которые клиент может видеть
		AllowCredentials: true, // Разрешить отправку куки и авторизационных заголовков
		MaxAge: 0, // Как долго кэшировать preflight-ответ (в секундах)
	})

	fmt.Println("Запуск на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}