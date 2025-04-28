package repositories

import (
	"BE_Ecommerce/src/entity"

	"gorm.io/gorm"
)

func GetAllCategories(db *gorm.DB) ([]entity.Category, error) {
	categories := []entity.Category{}
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
