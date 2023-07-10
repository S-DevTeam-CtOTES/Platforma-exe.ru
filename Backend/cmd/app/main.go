package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/avdushin/rgoauth/pkg/models/database"
	"github.com/avdushin/rgoauth/pkg/routes"
	"github.com/avdushin/rgoauth/vars"
)

func main() {
	// Установка соединения с базой данных MySQL
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("[DB] SQL adress: ", vars.DBConn, vars.DBName)
	if vars.PORT == "" {
		vars.PORT = ":4000"
	}
	/* Data Base */
	// Init DB
	database.CreateDB(vars.DBName)
	// Init DB tables
	database.InitTables()
	// Create Admins
	database.CreateAdmins()

	// Init router
	r := routes.SetupRoutes()

	/*
	* Start backup DB
	* Every 4 hours
	 */
	go func() {
		for range time.Tick(4 * time.Hour) {
			database.BackupDB()
			database.SaveStorage()
		}
	}()

	log.Printf("Server run at the %s%s", vars.DBHost, vars.PORT)

	// Start server
	log.Fatal(http.ListenAndServe(vars.PORT, r))

	// Start SSL (Production Server)
	// log.Fatal(http.ListenAndServeTLS(vars.PORT, vars.Cert, vars.Key, r))
}
