package models

import "golang.org/x/crypto/bcrypt"

type Password string

func NewPassword(val string) Password {
	return Password(val)
}

func (p Password) Hash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 8)
	return string(bytes), err
}

func (p Password) Check(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}

type User struct {
	FullName string   `json:"full_name" example:"Иван Иванов"`
	Username string   `json:"username" example:"ivan"`
	Password Password `json:"password,omitempty" example:"pass"`
}

type Users []User
