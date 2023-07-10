package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/avdushin/rgoauth/pkg/utils"
	"github.com/avdushin/rgoauth/vars"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding user registration request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Set DEFAULT `role` value
	if user.Role == "" {
		user.Role = "user"
	}

	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if user with the same username exists
	stmt, err := db.Prepare("SELECT id FROM users WHERE username = ?")
	if err != nil {
		log.Println("Не удалось выполнить запрос к БД:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(user.Username).Scan(&id)
	if err == nil {
		// User already exists
		http.Error(w, "Пользователь уже существует", http.StatusConflict)
		return
	}

	// Register a new user
	stmt, err = db.Prepare("INSERT INTO users (username, email, password, name, surname, patronymic, age, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Не удалось создать пользователя", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка хеширования пароля:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(user.Username, user.Email, string(hashedPassword), user.Name, user.Surname, user.Patronymic, user.Age, user.Role)

	if err != nil {
		log.Println("Ошибка добавления данных в БД:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"message": "Пользователь успешно зарегистрирован"}
	json.NewEncoder(w).Encode(response)

	log.Printf("User created: %s\n", user.Username)
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding user login request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Ошибка подключения к БД:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if the user exists and the password matches
	stmt, err := db.Prepare("SELECT id, username, password, role FROM users WHERE email = ?")
	if err != nil {
		log.Println("Error preparing database query:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var hashedPassword string
	if err = stmt.QueryRow(user.Email).Scan(&user.ID, &user.Username, &hashedPassword, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Неверный Email или Пароль", http.StatusBadRequest)
			log.Printf("Email не найден\nОшибка: %v", err)
			return
		}
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		log.Printf("Неверный пароль\nОшибка: %v", err)
		return
	}

	// Return the user's information in the response body
	response := struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Println("Выполнен вход под именем пользователя:", user.Username)
}
