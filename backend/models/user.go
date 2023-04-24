package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email,omitempty"`
	IsCurrentUser bool      `json:"isCurrentUser,omitempty"`
	IsFollower    bool      `json:"isFollower,omitempty"`
	IsFollowee    bool      `json:"isFollowee,omitempty"`
	Password      string    `json:"-"`
	PasswordHash  string    `json:"-"`
	LastLogin     time.Time `json:"lastLoginTime,omitempty"`
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
