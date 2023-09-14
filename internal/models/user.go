package models

import (
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint
	FullName  string `gorm:"type:varchar(256);not null"`
	Email     string `gorm:"type:varchar(256);unique;not null"`
	Active    bool   `gorm:"default:false;not null"`
	Password  string `gorm:"type:varchar(256);not null"`
	CreatedAt time.Time
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.FullName, validation.Required, validation.NewStringRule(func(fullName string) bool {
			names := strings.Split(fullName, " ")
			for _, name := range names {
				rule1 := govalidator.IsAlpha(name)
				rule2 := name != ""
				if !(rule1 && rule2) {
					return false
				}
			}
			return true
		}, "must be a valid full name")),
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
