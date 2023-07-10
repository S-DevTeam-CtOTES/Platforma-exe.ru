package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        string `json:"age"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
