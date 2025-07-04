package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"time"

	"github.com/Roman77St/simple_project/rest_api/models"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// sqlite

// func InitDatabase() error {
// 	var err error
// 	DB, err = sql.Open("sqlite3", "./sqlite3.db")
// 	if err != nil {
// 		return err
// 	}
// 	query := `CREATE TABLE IF NOT EXISTS users (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		age INTEGER,
// 		email TEXT UNIQUE,
// 		password TEXT DEFAULT ''
// 			);`
// 	_, err = DB.Exec(query)
// 	if err != nil {
// 		return err
// 	}
// fmt.Println("Соединение с базой данных установлено.")
// return nil
// 		}


// postgres

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
	query := `CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name varchar(50) NOT NULL,
			age INTEGER NOT NULL,
			email varchar(50) UNIQUE,
			password varchar(50) DEFAULT ''
			);`
	// query:= `DROP TABLE users` // для разработки, для изменения таблиц
		_, err = DB.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Соединение с базой данных установлено.")
	return nil
	}

func GetFromSQL(id string, user models.User) []byte {

	row := DB.QueryRow("SELECT * FROM users WHERE id = $1", id) // sqlite вместо $1 нужно ?
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
	}
	response, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
