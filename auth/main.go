package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
    _ "github.com/lib/pq"

)

type User struct {
	Email string  `json:"email"`
	Password string  `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	// Можете добавить другие данные, например, user ID, роль и т.д.
	UserEmail   string `json:"userEmail,omitempty"`
}

var DB *sql.DB

// Секретный ключ для подписи JWT-токена
var jwtSecretKey = []byte("very-secret-key")

func generateToken(user User) (string, error) {
    claims := jwt.MapClaims{
        "user": user.Email,
        "exp":  time.Now().Add(time.Hour * 24).Unix(), // Срок действия — 24 часа
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(jwtSecretKey)
}

func InitDatabase() (error) {
	var err error
	fmt.Println("Запускается приложение. соединение с базой данных через 5 секунд")
	time.Sleep(time.Second * 5) // для docker compose! Без задержки  нормально не запускается!
	connStr := "postgres://db_user:db_password@db:5432/db_name?sslmode=disable" // для compose!!!
	// connStr := "postgres://db_user:db_password@localhost:5433/db_name?sslmode=disable" // для go run!!!

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	fmt.Println("Соединение с базой данных установлено.")
	return nil
	}

func LoginUser(w http.ResponseWriter, request *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

    possiblePassword := user.Password
    row := DB.QueryRow("SELECT email, password FROM users WHERE email=$1;", user.Email) // sqlite вместо $1 нужно ?
    var user1 User
	err = row.Scan(&user1.Email, &user1.Password)
	if err != nil {
        fmt.Println(err)
	}
    validPassword := user1.Password
	if possiblePassword == validPassword {
		fmt.Println("Вход разрешен")
		token, err := generateToken(user1)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(AuthResponse{
			Token: token,
			UserEmail: user1.Email,
		})
	} else {
		fmt.Println("Неправильный логин или пароль")
	}
}

func main () {
    err := InitDatabase()
	if err != nil {
		log.Fatal("Ошибка подключения к реляционной базе данных: ", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/auth/login", LoginUser).Methods("POST", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
								 "http://127.0.0.1:8001",
								 "http://localhost:8001",
								 "https://strrv.ru",
								 "https://www.strrv.ru",
								},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешенные методы
		AllowedHeaders: []string{"*"}, // Разрешенные заголовки
		ExposedHeaders: []string{"Content-Length"}, // Заголовки, которые клиент может видеть
		AllowCredentials: true, // Разрешить отправку куки и авторизационных заголовков
		MaxAge: 300, // Как долго кэшировать preflight-ответ (в секундах)
	})

	fmt.Println("Запуск на порту 8082")
	log.Fatal(http.ListenAndServe(":8082", c.Handler(router)))
}