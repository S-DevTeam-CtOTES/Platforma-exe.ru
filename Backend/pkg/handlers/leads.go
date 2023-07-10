package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avdushin/rgoauth/pkg/utils"
	"github.com/avdushin/rgoauth/vars"
	_ "github.com/go-sql-driver/mysql"
)

type Leads struct {
	ID                       int    `json:"id"`
	LastName                 string `json:"lastName"`
	FirstName                string `json:"firstName"`
	Patronymic               string `json:"patronymic"`
	Birthday                 string `json:"birthday"`
	Phone                    string `json:"phone"`
	Email                    string `json:"email"`
	EducationLevel           string `json:"educationLevel"`
	AdditionalEducationLevel string `json:"additionalEducationLevel"`
	Direction                string `json:"direction"`
}

// Добавляем данные в таблицу лидов (далее импрот в Excel)
func AddDataToExcelFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг данных из тела запроса
	var Lead Leads
	if err := json.NewDecoder(r.Body).Decode(&Lead); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Подключение к базе данных MySQL
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Вставка данных в таблицу leads
	insertQuery := fmt.Sprintf("INSERT INTO leads (last_name, first_name, patronymic, birthday, phone, email, education_level, additional_education_level, direction) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		Lead.LastName, Lead.FirstName, Lead.Patronymic, Lead.Birthday, Lead.Phone, Lead.Email, Lead.EducationLevel, Lead.AdditionalEducationLevel, Lead.Direction)

	_, err = db.Exec(insertQuery)
	if err != nil {
		http.Error(w, "Failed to insert data into table", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	response := map[string]string{
		"message": "Data added successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// получаем все данные из БД лидов
func GetLeads(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.Query("SELECT id, last_name, first_name, patronymic, birthday, phone, email, education_level, additional_education_level, direction  FROM leads")
	if err != nil {
		log.Println("Error querying user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var leads []Leads
	for rows.Next() {
		var lead Leads
		err := rows.Scan(&lead.ID, &lead.LastName, &lead.FirstName, &lead.Patronymic, &lead.Birthday, &lead.Phone, &lead.Email, &lead.EducationLevel, &lead.AdditionalEducationLevel, &lead.Direction)
		if err != nil {
			log.Println("Error scanning user data:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		leads = append(leads, lead)
	}

	response, err := json.Marshal(leads)
	if err != nil {
		log.Println("Error encoding user data:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
