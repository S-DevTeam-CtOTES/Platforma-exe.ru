package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	h "github.com/avdushin/rgoauth/pkg/handlers"
)

// Setup routes...
func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// Обработчик OPTIONS запросов
	router.Methods(http.MethodOptions).HandlerFunc(handleOptions)

	/*
	* Handlers
	 */
	// Register
	router.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)
	// Login
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
	// Profile
	router.HandleFunc("/api/users/{id}", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", h.UpdateUser).Methods(http.MethodPut)
	// Admin
	router.HandleFunc("/admin/administrators", h.GetAdmins).Methods(http.MethodGet)
	router.HandleFunc("/api/users", h.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", h.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}/role", h.UpdateUserRole).Methods(http.MethodPut)
	// Courses
	router.HandleFunc("/courses", h.GetCourses).Methods("GET")
	router.HandleFunc("/courses", h.CreateCourse).Methods("POST")
	router.HandleFunc("/courses/{id}", h.GetCourse).Methods("GET")
	router.HandleFunc("/courses/{id}", h.UpdateCourse).Methods("PUT")
	router.HandleFunc("/courses/{id}", h.DeleteCourse).Methods("DELETE")
	// leads
	router.HandleFunc("/api/addDataToExcelForm", h.AddDataToExcelFormHandler).Methods("POST")
	router.HandleFunc("/api/getFormData", h.GetLeads).Methods("GET")

	// CORS settings
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return corsHandler(router)
}

// Handle options to allow signUP/Login
func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
