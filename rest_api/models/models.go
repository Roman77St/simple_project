package models

type User struct {
	ID string `json:"id" redis:"id"`
	Name string `json:"name" redis:"name"`
	Age int `json:"age" redis:"age"`
	Email string `json:"email" redis:"email"`
	Password string `json:"password" redis:"password"`
}
