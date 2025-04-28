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

func GetCategoryById(db *gorm.DB, categoryId int) (entity.Category, error) {
	category := entity.Category{}
	if err := db.Where("category_id = ?", categoryId).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func CreateCategory(db *gorm.DB, category *entity.CategoryRequest) error {
	if err := db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCategory(db *gorm.DB, categoryId int, category *entity.CategoryRequest) error {
	if err := db.Model(&entity.Category{}).Where("category_id = ?", categoryId).Updates(category).Error; err != nil {
		return err
	}
	return nil
}
