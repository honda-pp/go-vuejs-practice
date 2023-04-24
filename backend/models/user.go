package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int
	Username      string
	Email         string    `json:",omitempty"`
	Password      string    `json:"-"`
	PasswordHash  string    `json:"-"`
	LastLoginTime time.Time `json:"-"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
