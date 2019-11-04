package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	userFreelancer = "freelancer"
	userCustomer   = "client"
	)

type User struct {
	ID 				int64 `json:"id" valid:"int, optional"`
	FirstName 		string `json:"firstName" valid:"utfletter, required"`
	SecondName 		string `json:"secondName" valid:"utfletter"`
	UserName     	string `json:"username" valid:"alphanum"`
	Email 			string `json:"email" valid:"email"`
	Password		string `json:"password" valid:"length(6|100)"`
	EncryptPassword string `json:"-" valid:"-"`
	Avatar 			[]byte `json:"-" valid:"-"`
	UserType 		string `json:"type" valid:"in(client|freelancer)"`
}


func (u *User) BeforeCreate() error {
	if len(u.UserType) == 0 || u.UserType != userFreelancer && u.UserType != userCustomer {
		u.UserType = userFreelancer
	}

	if len(u.Password) > 0 {
		enc, err := EncryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptPassword = enc
	}
	u.Password = ""
	return nil
}

func (u *User) SetUserType(userType string) error {
	if userType == userFreelancer || userType == userCustomer {
		u.UserType = userType
		return nil
	}
	return errors.New("wrong user type")
}

func (u *User) IsManager() bool {
	return u.UserType == userCustomer
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptPassword), []byte(password)) == nil
}

func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//func (u *User) Validate() error {
//	return validation.ValidateStruct(
//		u,
//		validation.Field(&u.Email, validation.Required, is.Email),
//		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
//	)
//}
//
//func requiredIf(cond bool) validation.RuleFunc {
//	return func(value interface{}) error {
//		if cond {
//			return validation.Validate(value, validation.Required)
//		}
//		return nil
//	}
//}
//
