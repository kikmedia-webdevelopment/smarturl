package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/juliankoehn/mchurl/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
	Email        string     `json:"email" gorm:"unique;not null"`
	Password     string     `json:"password,omitempty"`
	Token        string     `json:"token,omitempty"`
	TokenExpires time.Time  `json:"token_expires"`
}

func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	if u.Password == "" {
		pass, err := utils.RandomPass(12)
		if err != nil {
			return err
		}
		u.Password = pass
	}
	pw, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pw
	return
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}
