package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetUser(db *sql.DB, username string) (User, error) {
	var user User
	rows, err := db.Query("SELECT id, username, email, password_hash FROM users WHERE username = $1", username)
	defer rows.Close()
	if err != nil {
		return user, err
	}
	if rows.Next() {
		var id int
		var username string
		var email string
		var passwordHash string
		if err := rows.Scan(&id, &username, &email, &passwordHash); err != nil {
			return user, err
		}
		user = User{
			ID:           id,
			Username:     username,
			Email:        email,
			PasswordHash: passwordHash,
		}
	}
	return user, nil
}
