package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
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
	)
}

func (u *User) HashPassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	u.Password = string(hashedBytes)
	return nil
}

func (u *User) ComparePassword(compareTo string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(compareTo))
	return err == nil
}
