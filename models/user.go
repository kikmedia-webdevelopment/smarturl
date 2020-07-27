package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/juliankoehn/mchurl/utils"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User is a user of our system
type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
	Email        string     `json:"email" gorm:"unique;not null"`
	Password     string     `json:"password,omitempty"`
	Token        string     `json:"token,omitempty"`
	TokenExpires time.Time  `json:"token_expires"`

	RecoveryToken  string     `json:"-" db:"recovery_token"`
	RecoverySentAt *time.Time `json:"recovery_sent_at,omitempty" db:"recovery_sent_at"`

	EmailChangeToken  string     `json:"-" db:"email_change_token"`
	EmailChange       string     `json:"new_email,omitempty" db:"email_change"`
	EmailChangeSentAt *time.Time `json:"email_change_sent_at,omitempty" db:"email_change_sent_at"`
	LastSignInAt      *time.Time `json:"last_sign_in_at,omitempty" db:"last_sign_in_at"`
}

// SetEmail updates the email of the User
func (u *User) SetEmail(tx *storage.Connection, email string) error {
	u.Email = email
	return tx.Model(u).Update("email", email).Error
}

// UpdatePassword updates the given password
func (u *User) UpdatePassword(tx *storage.Connection, password string) error {
	pw, err := hashPassword(password)
	if err != nil {
		return nil
	}
	u.Password = pw
	return tx.Model(u).Update("password", pw).Error
}

func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// ConfirmEmailChange confirm the change of email for a user
func (u *User) ConfirmEmailChange(tx *storage.Connection) error {
	u.Email = u.EmailChange
	u.EmailChange = ""
	u.EmailChangeToken = ""
	return tx.Model(u).Updates(User{Email: u.EmailChange, EmailChange: "", EmailChangeToken: ""}).Error
}

// Recover resets the recovery token
func (u *User) Recover(tx *storage.Connection) error {
	u.RecoveryToken = ""
	return tx.Model(u).Update("recovery_token", "").Error
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

func (u *User) BeforeSave(scope *gorm.Scope) error {
	if u.RecoverySentAt != nil && u.RecoverySentAt.IsZero() {
		u.RecoverySentAt = nil
	}

	if u.EmailChangeSentAt != nil && u.EmailChangeSentAt.IsZero() {
		u.EmailChangeSentAt = nil
	}
	if u.LastSignInAt != nil && u.LastSignInAt.IsZero() {
		u.LastSignInAt = nil
	}

	return nil
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

func findUser(tx *storage.Connection, query string, args ...interface{}) (*User, error) {
	obj := &User{}
	if err := tx.Where(query, args...).First(obj).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, UserNotFoundError{}
		}
		return nil, errors.Wrap(err, "error finding user")
	}
	return obj, nil
}

// FindUserByEmail finds a user with the matching email.
func FindUserByEmail(tx *storage.Connection, email string) (*User, error) {
	return findUser(tx, "email = ?", email)
}

// FindUserByID finds a user matching the provided ID.
func FindUserByID(tx *storage.Connection, id uint) (*User, error) {
	return findUser(tx, "id = ?", id)
}

// FindUsers finds users.
func FindUsers(tx *storage.Connection) ([]*User, error) {

	users := make([]*User, 0)
	if err := tx.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// IsDuplicatedEmail returns whether a user exists with a matching email.
func IsDuplicatedEmail(tx *storage.Connection, email string) (bool, error) {
	_, err := FindUserByEmail(tx, email)
	if err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
