package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/avdushin/rgoauth/pkg/utils"
	"github.com/avdushin/rgoauth/vars"
)

func GetAdmins(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.Query("SELECT id, username, email, name, surname, patronymic, age, role FROM users WHERE role = 'admin'")
	if err != nil {
		log.Println("Error querying user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []User
	for rows.Next() {
		var admin User
		err := rows.Scan(&admin.ID, &admin.Username, &admin.Email, &admin.Name, &admin.Surname, &admin.Patronymic, &admin.Age, &admin.Role)
		if err != nil {
			log.Println("Error scanning user data:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		admins = append(admins, admin)
	}

	response, err := json.Marshal(admins)
	if err != nil {
		log.Println("Error encoding user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
