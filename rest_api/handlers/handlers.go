package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Roman77St/simple_project/rest_api/models"
	"github.com/Roman77St/simple_project/storage"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// Список разрешенных источников

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	redisKey := "list_user"
	res, err := storage.RedisDB.Get(redisKey).Result()
	if err == redis.Nil {
		fmt.Println("Значение в Redis не установлено")
		// Взять значение из sql-базы
		rows, err := storage.DB.Query("SELECT * FROM users")
		if err != nil {
			fmt.Println("ОШИБКА!", err)
		}
		defer rows.Close()
		users := []models.User{}
		for rows.Next(){
			u := models.User{}
			err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Email, &u.Password)
			if err != nil{
				fmt.Println("Ошибка в функции GetUsers", err)
				continue
			}
			users = append(users, u)
		}
		allUsers, _ := json.Marshal(users)
		storage.SetToRedis(redisKey, allUsers)
		fmt.Fprint(w, string(allUsers))
		return
	} else if err != nil {
		fmt.Printf("Ошибка при получении JSON из Redis: %v\n", err)
	}
	fmt.Println("Взято кешированное значение из Redis")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)

}

func GetUser(w http.ResponseWriter, request *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	userID := params["id"]
	redisKey := "user:" + userID
	res, err := storage.RedisDB.Get(redisKey).Result()
	if err == redis.Nil {
		fmt.Println("Значение в Redis не установлено")
		// Взять значение из sql-базы, return
		response := storage.GetFromSQL(userID, user)
		storage.SetToRedis(redisKey, response)
		fmt.Fprint(w, string(response))
		return
	} else if err != nil {
		fmt.Printf("Ошибка при получении JSON из Redis: %v\n", err)
	}
	fmt.Println("Взято кешированное значение из Redis")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)

}


func CreateUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name)
	_, err = storage.DB.Exec("INSERT INTO users (name, age, email, password) VALUES ($1, $2, $3, $4);", user.Name, user.Age, user.Email, user.Password)  // sqlite вместо $1, $2 нужно ?, ?
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUser (w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, _ := strconv.Atoi(params["id"])
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	_, err = storage.DB.Exec("UPDATE users SET name=$1, age=$2, email=$3, password=$4 where id=$5;", user.Name, user.Age, user.Email, user.Password, userID) // sqlite вместо $1, $2, $3 нужно ?, ?, ?
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteUser (w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, _ := strconv.Atoi(params["id"])
	_, err := storage.DB.Exec("DELETE FROM users WHERE id=$1;", userID)  // sqlite вместо $1 нужно ?
	if err != nil {
		fmt.Println(err)
	}
}