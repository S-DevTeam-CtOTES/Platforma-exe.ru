package database

import (
	"archive/zip"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"github.com/avdushin/rgoauth/pkg/handlers"
	"github.com/avdushin/rgoauth/vars"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

/* work with MySQL DataBase */

// Make querry...
func DBQuerry(querry, comment string) {
	// DBConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)

	// Соедененеие с базой данных
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // закрываем подключение к БД

	// Делаем запрос
	if _, err = db.Exec(fmt.Sprint(querry)); err != nil {
		log.Fatal(err)
	}

	// Выводим лог
	log.Println(comment)
}

// Создаем БД
func CreateDB(name string) {
	// CREATE DATABASE IF NOT EXISTS
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // закрываем подключение к БД

	// Делаем запрос
	if _, err = db.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, name)); err != nil {
		log.Fatal(err)
	}

	// Выводим лог
	log.Printf("[DB] Соединение с БД %s успешно установлено...", name)
}

// Update user's roles
func UpdateRoles() {
	// UPDATE users SET role = 'user' WHERE role = '';
	DBQuerry(`
	UPDATE users SET role = 'user' WHERE role = '';
`, "[DB] Роли пользователей обновлены")
}

// Создаем нужные таблицы
func InitTables() {
	// Create users table
	DBQuerry(`
	CREATE TABLE IF NOT EXISTS users (
		id INT NOT NULL AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL DEFAULT '',
		email VARCHAR(50) NOT NULL DEFAULT '',
		password VARCHAR(255) NOT NULL DEFAULT '',
		name VARCHAR(255) DEFAULT '',
		surname VARCHAR(255) DEFAULT '',
		patronymic VARCHAR(255) DEFAULT '',
		age INT UNSIGNED DEFAULT 0,
		role VARCHAR(255) DEFAULT 'user',
		PRIMARY KEY (id),
		UNIQUE KEY (username),
		UNIQUE KEY (email)
	  );`, "[DB] Таблица users: OK...")
	// Create courses table
	DBQuerry(`
	CREATE TABLE IF NOT EXISTS courses (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		description TEXT(5000) NOT NULL,
		price INT NOT NULL,
		img TEXT(3000) NOT NULL,
		registration_link VARCHAR(255) NOT NULL,
		page VARCHAR(255) NOT NULL,
		time VARCHAR(255) NOT NULL,
		type VARCHAR(255) NOT NULL,
		type_event VARCHAR(255) NOT NULL,
		video_url VARCHAR(255)
	);
	`, "[DB] Таблица courses: OK...")

	// Create leads table
	DBQuerry(`
    CREATE TABLE IF NOT EXISTS leads (
        id INT AUTO_INCREMENT PRIMARY KEY,
        last_name VARCHAR(255) NOT NULL,
        first_name VARCHAR(255) NOT NULL,
        patronymic VARCHAR(255),
        birthday VARCHAR(255),
        phone VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        education_level TEXT,
        additional_education_level TEXT,
        direction TEXT
    );    
    `, "[DB] Таблица leads: OK...")
}

// Create Admin users
func CreateAdmins() {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Проверка наличия пользователя admin
	adminExists := false
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", vars.ADMIN_NAME).Scan(&adminExists)
	if err != nil {
		log.Fatal("Ошибка при проверке наличия пользователя:", err)
	}

	if !adminExists {
		// Хеширование пароля
		password := vars.ADMIN_PASSWORD
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Ошибка при хешировании пароля:", err)
		}

		// Создание пользователя admin
		admin := handlers.User{
			Username: vars.ADMIN_NAME,
			Email:    vars.ADMIN_EMAIL,
			Password: string(hashedPassword),
			Role:     "admin",
		}

		// Вставка пользователя в базу данных
		_, err = db.Exec("INSERT INTO users (username, email, password, name, surname, patronymic, age, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			admin.Username, admin.Email, admin.Password, admin.Name, admin.Surname, admin.Patronymic, admin.Age, admin.Role)
		if err != nil {
			log.Fatal("Ошибка при создании пользователя:", err)
		}

		log.Printf("[DB] Пользователь %s успешно создан!", vars.ADMIN_NAME)
	} else {
		log.Printf("[DB] Пользователь %s уже существует!", vars.ADMIN_NAME)
	}
}

// Создаём бэкап БД
func BackupDB() {
	// Получаем текущее время
	now := time.Now().Local()
	DateFormat := now.Format("2006-01-02_15-04-05")

	// Создаём папку для хранения бэкапов если её ещё не существует
	if _, err := os.Stat("./backups"); os.IsNotExist(err) {
		os.Mkdir("./backups", os.ModePerm)
	}
	// Создаём папку для хранения архивов бэкапов если её ещё не существует
	if _, err := os.Stat("./backups/archives"); os.IsNotExist(err) {
		os.Mkdir("./backups/archives", os.ModePerm)
	}

	// Имя файла для бэкапа и zip архива
	backupFileName := fmt.Sprintf("./backups/%s.sql", DateFormat)
	// zipFileName := fmt.Sprintf("./backups/archives/%s.zip", DateFormat)

	// Создаём бэкап
	backupFile, err := os.Create(backupFileName)
	if err != nil {
		log.Fatalf("[BACKUP] Ошибка создания бэкап файла БД: %v", err)
	}
	defer backupFile.Close()

	// mysqldump command
	cmdArgs := []string{
		"-u",
		vars.DBUser,
		"-p" + vars.DBPass,
		"-h",
		vars.DBHost,
		vars.DBName,
	}

	// Бэкапим БД с помощью mysqldump
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// Windows
		cmd = exec.Command("C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin\\mysqldump.exe", cmdArgs...)
	default:
		// Unix
		cmd = exec.Command("mysqldump", cmdArgs...)
	}

	outfile, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating mysqldump command: %v", err)
	}

	// Запускаем mysqldump command
	if err := cmd.Start(); err != nil {
		log.Fatalf("Неудалось запустить mysqldump: %v", err)
	}

	// Записываем вывод mysqldump в бэкап файл
	if _, err := io.Copy(backupFile, outfile); err != nil {
		log.Fatalf("Ошибка записи бэкап файла: %v", err)
	}

	// Ждём выполнения mysqldump
	if err := cmd.Wait(); err != nil {
		log.Fatalf("[Ошибка] Я устал ждать пока mysqldump соберет там свои дампы: %v", err)
	}

	log.Printf("[DB|BACKUP]: Резервная копия создана и сохранена в %s\n", backupFileName)

	// // Проверяем размер файла
	// fileInfo, err := os.Stat(backupFileName)
	// if err != nil {
	// 	log.Fatalf("Ошибка при получении информации о файле %s: %v", backupFileName, err)
	// }

	// if fileInfo.Size() > 4*1024*1024 { // 20 мб в байтах 20*1024*1024
	// 	// Создаем новый архив
	// 	zipFile, err := os.Create(zipFileName)
	// 	if err != nil {
	// 		log.Fatalf("Ошибка при создании zip файла %s: %v", zipFileName, err)
	// 	}
	// 	defer zipFile.Close()

	// 	zipWriter := zip.NewWriter(zipFile)
	// 	defer zipWriter.Close()

	// 	// Добавляем файл в архив
	// 	err = addFileToZip(backupFileName, zipWriter)
	// 	if err != nil {
	// 		log.Fatalf("Ошибка добавления файла %s в zip архив: %v", backupFileName, err)
	// 	}
	// 	fmt.Printf("Архив %s успешно создан!\n", zipFileName)
	// } else {
	// 	fmt.Printf("Файл %s не нуждается в архивации!\n", backupFileName)
	// }
}

// Функция для добавления файла в архив
func addFileToZip(filePath string, zipWriter *zip.Writer) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Получаем информацию о файле
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Создаем новый файл в архиве
	zipFile, err := zipWriter.Create(fileInfo.Name())
	if err != nil {
		return err
	}

	// Копируем содержимое файла в архив
	_, err = io.Copy(zipFile, file)
	if err != nil {
		return err
	}

	return nil
}

func SaveStorage() {
	// Получаем текущее время
	now := time.Now().Local()
	DateFormat := now.Format("2006-01-02_15-04-05")
	zipFileName := fmt.Sprintf("./backups/archives/%s.zip", DateFormat)

	// Получаем список всех файлов в папке с расширением .sql
	fileList, err := filepath.Glob(filepath.Join("./backups", "*.sql"))
	if err != nil {
		log.Fatalf("Ошибка получения списка файлов: %v", err)
	}

	// Проверяем, есть ли больше 5 файлов в папке
	if len(fileList) > 5 {
		// Сортируем список файлов по времени изменения (последний файл будет первым в списке)
		sort.Slice(fileList, func(i, j int) bool {
			fileStatI, err := os.Stat(fileList[i])
			if err != nil {
				log.Fatalf("Ошибка получения информации о файле %s: %v", fileList[i], err)
			}
			fileStatJ, err := os.Stat(fileList[j])
			if err != nil {
				log.Fatalf("Ошибка получения информации о файле %s: %v", fileList[j], err)
			}
			return fileStatI.ModTime().After(fileStatJ.ModTime())
		})

		// Создаем новый архив
		// zipFileName := fmt.Sprintf("%s.zip", fileList[0])
		zipFile, err := os.Create(zipFileName)
		if err != nil {
			log.Fatalf("Ошибка создания zip файла %s: %v", zipFileName, err)
		}
		defer zipFile.Close()

		zipWriter := zip.NewWriter(zipFile)

		// Добавляем последний файл в архив
		err = addFileToZip(fileList[0], zipWriter)
		if err != nil {
			log.Fatalf("Ошибка добавления файла %s в zip архив: %v", fileList[0], err)
		}

		// Закрываем архив
		err = zipWriter.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия zip архива: %v", err)
		}

		// Удаляем все файлы .sql
		for _, file := range fileList[1:] {
			err = os.Remove(file)
			if err != nil {
				log.Fatalf("Ошибка удаления файла %s: %v", file, err)
			}
		}

		fmt.Printf("[BACKUP] Файл %s успешно заархивирован и удалены все файлы .sql в папке!\n", fileList[0])
	}
}
