package vars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// app
	VERSION = "1.2"
	_       = godotenv.Load()
	// .env vars
	PORT = os.Getenv("PORT")
	// Получение переменных окружения
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	// Admin
	ADMIN_NAME     = os.Getenv("ADMIN_NAME")
	ADMIN_EMAIL    = os.Getenv("ADMIN_EMAIL")
	ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
	// DB connection...
	DBConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", DBUser, DBPass, DBHost, DBPort)
	// SSL/TLS certs path
	Cert = "/var/www/certs/domain-example.ru.pub"
	Key  = "/var/www/private/domain-example.ru.key"
)
