package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	ID 				int64 `json:"id"`
	FirstName 		string `json:"firstName"`
	SecondName 		string `json:"secondName"`
	UserName     	string `json:"username"`
	Email 			string `json:"email"`
	Password 		string `json:"password, omitempty"`
	EncryptPassword string `json:"-"`
	Avatar string `json:"-"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptPassword = enc
	}
	u.Password = ""
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

type UserInput struct {
	Name 		string `json:"name"`
	Surname 	string `json:"surname"`
	Email   	string `json:"email"`
	Password	string `json:"password"`
}

func (u *UserInput) CheckEmail() error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if valid := re.MatchString(u.Email); valid == false {
		return errors.New("invalid email")
	}
	return nil
}
