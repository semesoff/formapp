package data

import (
	"database/sql"
	"fmt"
	"formapp/internal/config"
	"formapp/internal/model"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		return nil, err
	}
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
		return nil, err
	}
	log.Printf("Database is connected")
	return db, nil
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100),
        number VARCHAR(20),
        address TEXT,
        education TEXT,
        experience TEXT,
        skills TEXT
    );`
	_, err := db.Exec(query)
	return err
}

func InsertData(db *sql.DB, userData model.UserData) error {
	query := `INSERT INTO users (name, email, number, address, education, experience, skills) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(query, userData.Name, userData.Email, userData.Number, userData.Address, userData.Education, userData.Experience, userData.Skills)
	return err
}

func GetAllUserData(db *sql.DB) ([]model.UserData, error) {
	rows, err := db.Query("SELECT name, email, number, address, education, experience, skills FROM users")
	if err != nil {
		log.Fatalf("Error getting values from database: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []model.UserData
	for rows.Next() {
		var user model.UserData
		err := rows.Scan(&user.Name, &user.Email, &user.Number, &user.Address, &user.Education, &user.Experience, &user.Skills)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
