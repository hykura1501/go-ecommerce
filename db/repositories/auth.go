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

func GetUserByProvider(dbInstance *gorm.DB, loginProvider int, providerId string) (*entity.User, error) {
	var user entity.User
	result := dbInstance.Where("provider_id = ? AND login_provider = ?", providerId, loginProvider).First(&user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUserByProvider(dbInstance *gorm.DB, user *entity.User) error {
	err := dbInstance.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
