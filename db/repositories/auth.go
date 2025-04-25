package repositories

import (
	"BE_Ecommerce/src/entity"

	"gorm.io/gorm"
)

func GetUserByUsername(dbInstance *gorm.DB, username string) (*entity.User, error) {
	var user entity.User
	err := dbInstance.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create and return user
func CreateUser(dbInstance *gorm.DB, user *entity.User) (*entity.User, error) {
	err := dbInstance.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
