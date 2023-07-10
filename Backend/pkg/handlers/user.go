package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/avdushin/rgoauth/pkg/utils"
	"github.com/avdushin/rgoauth/vars"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// получаем данные пользователя
func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]

	var user User
	err = db.QueryRow("SELECT id, username, email, name, surname, patronymic, age, role FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Role)
	if err != nil {
		log.Println("Error querying user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println("Error encoding user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Обновляем данные пользователя
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == http.MethodPut {
		db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
		if err != nil {
			log.Println("Ошибка подклбчения к Базе Данных", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		id := mux.Vars(r)["id"]

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Error decoding user update request:", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Хеширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Ошибка хеширования пароля:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		stmt, err := db.Prepare("UPDATE users SET username=?, email=?, password=?, name=?, surname=?, patronymic=?, role=?, age=? WHERE id=?")
		if err != nil {
			log.Println("Ошибка обновления данных пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		if _, err = stmt.Exec(user.Username, user.Email, string(hashedPassword), user.Name, user.Surname, user.Patronymic, user.Role, user.Age, id); err != nil {
			log.Println("Ошибка обновления данных пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"message": "Данные успешно обновлены",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Обновляем роль пользователя
func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == http.MethodPut {
		db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
		if err != nil {
			log.Println("Ошибка подклбчения к Базе Данных", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		id := mux.Vars(r)["id"]

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Error decoding user update request:", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("UPDATE users SET role=? WHERE id=?")
		if err != nil {
			log.Println("Ошибка обновления роли пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		if _, err = stmt.Exec(user.Role, id); err != nil {
			log.Println("Ошибка обновления роли пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"message": "Роль пользователя успешно обновлена",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Получаем всех пользователей
func GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, username, email, name, surname, patronymic, age, role FROM users")
	if err != nil {
		log.Println("Error querying user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Role)
		if err != nil {
			log.Println("Error scanning user data:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	response, err := json.Marshal(users)
	if err != nil {
		log.Println("Error encoding user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Добавляем пользователя
func AddUser(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == http.MethodPost {
		db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
		if err != nil {
			log.Println("Ошибка подключения к Базе Данных", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Error decoding user add request:", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Хеширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Ошибка хеширования пароля:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		stmt, err := db.Prepare("INSERT INTO users (username, email, password, name, surname, patronymic, role, age) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Println("Ошибка добавления пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.Username, user.Email, string(hashedPassword), user.Name, user.Surname, user.Patronymic, user.Role, user.Age)
		if err != nil {
			log.Println("Ошибка добавления пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"message": "Пользователь успешно добавлен",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Удаляем пользователя
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == http.MethodDelete {
		db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
		if err != nil {
			log.Println("Ошибка подключения к Базе Данных", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		id := mux.Vars(r)["id"]

		stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
		if err != nil {
			log.Println("Ошибка удаления пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(id)
		if err != nil {
			log.Println("Ошибка удаления пользователя:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"message": "Пользователь успешно удален",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
