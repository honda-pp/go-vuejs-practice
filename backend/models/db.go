package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
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

	return &DB{db}, nil
}

func (db *DB) GetUserFromUsername(username string) (*User, error) {
	var user User
	row := db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", username)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetUserFromID(id string) (*User, error) {
	var user User
	row := db.QueryRow("SELECT id, username FROM users WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetUserList() ([]*User, error) {
	rows, err := db.Query(`
		SELECT
			users.id,
			users.username,
			MAX(login_history.login_time) AS last_login_time
		FROM
			users
		LEFT JOIN
			login_history ON users.id = login_history.user_id
		GROUP BY
			users.id
		ORDER BY
			last_login_time;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		var loginTime sql.NullTime
		err = rows.Scan(&user.ID, &user.Username, &loginTime)
		if err != nil {
			return nil, err
		}
		user.LastLoginTime = loginTime.Time
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (db *DB) AddUser(user User) error {
	_, err := db.Exec("INSERT INTO users(username, email, password_hash) VALUES($1, $2, $3)",
		user.Username, user.Email, user.PasswordHash)
	return err
}

func (db *DB) InsertLoginHistory(user *User) error {
	_, err := db.Exec("INSERT INTO login_history(user_id) VALUES($1)", user.ID)
	return err
}
