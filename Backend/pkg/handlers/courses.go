package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/avdushin/rgoauth/vars"
	"github.com/gorilla/mux"
)

type Course struct {
	ID               int    `json:"id"`                // course ID
	Name             string `json:"name"`              // course name
	Description      string `json:"description"`       // course description
	Price            int    `json:"price"`             // course price
	Img              string `json:"img"`               // course image (link to image url)
	VideoURL         string `json:"video_url"`         // course video URL
	RegistrationLink string `json:"registration_link"` // registration link to course (example: https://www.teta.mts.ru/course_name)
	Page             string `json:"page"`              // course page `/courses/:id`
	Time             string `json:"time"`              // course time (length, example: "4 месяца")
	Type             string `json:"type"`              // course type (example: "ДПО", "Практика", "Секции")
	TypeEvent        string `json:"type_event"`        // course type event (example: "Онлайн", "Оффлайн")
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	courses := []Course{}

	rows, err := db.Query("SELECT id, name, description, price, img, video_url, registration_link, page, time, type, type_event FROM courses")
	if err != nil {
		log.Println("Error querying courses:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Price, &course.Img, &course.VideoURL, &course.RegistrationLink, &course.Page, &course.Time, &course.Type, &course.TypeEvent)
		if err != nil {
			log.Println("Error scanning course row:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		courses = append(courses, course)
	}

	response, err := json.Marshal(courses)
	if err != nil {
		log.Println("Error encoding courses:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	course := Course{}

	row := db.QueryRow("SELECT id, name, description, price, img, video_url, registration_link, page, time, type, type_event FROM courses WHERE id = ?", id)
	err = row.Scan(&course.ID, &course.Name, &course.Description, &course.Price, &course.Img, &course.VideoURL, &course.RegistrationLink, &course.Page, &course.Time, &course.Type, &course.TypeEvent)
	if err != nil {
		log.Println("Error querying course:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(course)
	if err != nil {
		log.Println("Error encoding course:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var course Course
	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		log.Println("Error decoding course:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO courses (name, description, price, img, video_url, registration_link, page, time, type, type_event) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		course.Name, course.Description, course.Price, course.Img, course.VideoURL, course.RegistrationLink, course.Page, course.Time, course.Type, course.TypeEvent)
	if err != nil {
		log.Println("Error inserting course:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()

	response, err := json.Marshal(map[string]int64{"id": id})
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var course Course
	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		log.Println("Error decoding course:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE courses SET name = ?, description = ?, price = ?, img = ?, video_url = ?, registration_link = ?, page = ?, time = ?, type = ?, type_event = ? WHERE id = ?",
		course.Name, course.Description, course.Price, course.Img, course.VideoURL, course.RegistrationLink, course.Page, course.Time, course.Type, course.TypeEvent, id)
	if err != nil {
		log.Println("Error updating course:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	_, err = db.Exec("DELETE FROM courses WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting course:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
