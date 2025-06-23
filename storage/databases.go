package storage

import (
	"database/sql"
	"fmt"
	"time"

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
// 			name TEXT NOT NULL,
// 			age INTEGER
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
			age INTEGER NOT NULL
			);`
		_, err = DB.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Соединение с базой данных установлено.")
	return nil
	}
