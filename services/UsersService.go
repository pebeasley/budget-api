package services

import (
	"github.com/pebeasley/leaves-api/database"
	"github.com/pebeasley/leaves-api/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUserById(id string) (models.User, error) {
	var FoundUser models.User
	result := database.DB.
		Model(models.User{}).
		Where(&models.User{ID: id}).
		Find(&FoundUser)

	return FoundUser, result.Error
}

func GetUserByEmail(email string) (models.User, error) {
	var FoundUser models.User
	result := database.DB.
		Where(&models.User{Email: email}).
		Find(&FoundUser)

	return FoundUser, result.Error
}

func CreateNewUser(user *models.User) (*models.User, error) {
	hash, err := hashPassword(user.Password)
	if err != nil {
		return nil, nil
	}
	user.Password = hash
	result := database.DB.Create(&user)

	return user, result.Error
}
