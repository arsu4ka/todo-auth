package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID       uint
	FullName string
	Email    string `gorm:"unique"`
	Password string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.FullName, validation.Required),
		validation.Field(&u.Password, validation.Length(6, 25), is.Alphanumeric),
	)
}

func (u *User) HasPassword(password string) bool {
	return u.Password == password
}
