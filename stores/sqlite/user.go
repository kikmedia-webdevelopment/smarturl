package sqlite

import (
	"fmt"
	"time"

	"github.com/juliankoehn/mchurl/models"
)

func (s *SqliteStore) CreateUser(user models.User) (*models.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqliteStore) UserUpdateToken(id uint, token string) error {

	expires := time.Now().Add(time.Hour * 72)
	fmt.Println(expires)
	if err := s.db.
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"token": token, "token_expires": expires}).
		Error; err != nil {
		return err
	}
	return nil
}

func (s *SqliteStore) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqliteStore) FindUserByToken(token string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
