package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Roman77St/simple_project/storage"
	"github.com/gorilla/mux"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := storage.DB.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("ОШИБКА!", err)
	}
	defer rows.Close()
	users := []User{}
	for rows.Next(){
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
        if err != nil{
            fmt.Println(err)
            continue
		}
		users = append(users, u)
	}
	w.Header().Set("Content-Type", "application/json")
	allUsers, _ := json.Marshal(users)
	fmt.Fprint(w, string(allUsers))

}

func GetUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	userID, _ := strconv.Atoi(params["id"])
	row := storage.DB.QueryRow("SELECT * FROM users WHERE id = $1", userID) // sqlite вместо $1 нужно ?
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := json.Marshal(user)
	fmt.Fprint(w, string(response))
}

func CreateUser(w http.ResponseWriter, request *http.Request) {
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name)
	_, err = storage.DB.Exec("INSERT INTO users (name, age) VALUES ($1, $2);", user.Name, user.Age)  // sqlite вместо $1, $2 нужно ?, ?
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUser (w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, _ := strconv.Atoi(params["id"])
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	_, err = storage.DB.Exec("UPDATE users SET name=$1, age=$2 where id=$3;", user.Name, user.Age, userID) // sqlite вместо $1, $2, $3 нужно ?, ?, ?
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