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

func DeleteCategory(db *gorm.DB, categoryId int) error {
	if err := db.Where("category_id = ?", categoryId).Delete(&entity.Category{}).Error; err != nil {
		return err
	}
	return nil
}

func GetProductsByCategoryId(db *gorm.DB, categoryId int, paging *entity.Paging) ([]entity.Product, error) {
	products := []entity.Product{}
	querySQL := `
		SELECT 
			p.product_id, 
			p.product_name, 
			p.price, 
			p.stock, 
			p.description, 
			p.discount, 
			p.tag,
			jsonb_build_object('category_id', c.category_id, 'category_name', c.category_name) AS category,
			jsonb_build_object('manufacturer_id', m.manufacturer_id, 'manufacturer_name', m.manufacturer_name) AS manufacturer,
			COALESCE(jsonb_agg(pi.image_url) FILTER (WHERE pi.image_url IS NOT NULL), '[]'::jsonb) AS images
		FROM product p 
		LEFT JOIN product_image pi ON pi.product_id = p.product_id
		LEFT JOIN category c ON c.category_id = p.category_id
		LEFT JOIN manufacturer m ON m.manufacturer_id = p.manufacturer_id
		WHERE p.category_id = @category_id
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.category_name, 
			m.manufacturer_id, m.manufacturer_name
		ORDER BY p.created_at DESC
		LIMIT @limit OFFSET @offset;
	`

	err := db.Raw(querySQL, map[string]interface{}{
		"category_id": categoryId,
		"limit":       paging.Limit,
		"offset":      paging.Offset,
	}).Scan(&products).Error

	if err != nil {
		return nil, err
	}

	var count int64
	err = db.Model(&entity.Product{}).Where("category_id = ?", categoryId).Count(&count).Error
	if err != nil {
		return nil, err
	}

	paging.SetTotalPages(count)

	return products, nil
}
