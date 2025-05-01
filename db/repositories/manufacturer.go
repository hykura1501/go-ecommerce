package repositories

import (
	"BE_Ecommerce/src/entity"

	"gorm.io/gorm"
)

func GetAllManufacturers(dbInstance *gorm.DB) ([]entity.Manufacturer, error) {
	manufacturers := []entity.Manufacturer{}
	if err := dbInstance.Find(&manufacturers).Error; err != nil {
		return nil, err
	}
	return manufacturers, nil
}
