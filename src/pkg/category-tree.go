package pkg

import "BE_Ecommerce/src/entity"

func GetCategoryTree(categories []entity.Category, superCategoryId int) []entity.Category {
	categoryTree := []entity.Category{}

	for _, category := range categories {
		if category.SuperCategoryId == superCategoryId {
			category.Children = GetCategoryTree(categories, category.CategoryId)
			categoryTree = append(categoryTree, category)
		}
	}

	return categoryTree
}
