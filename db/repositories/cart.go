package repositories

import (
	"BE_Ecommerce/src/entity"

	"gorm.io/gorm"
)

func GetCartByUserId(dbInstance *gorm.DB, userId int, paging *entity.Paging) ([]entity.Cart, error) {
	var carts []entity.Cart
	query := `
		SELECT 
			json_build_object (
				'id', c.product_id,
				'name', p.product_name,
				'price', p.price,
				'discount', p.discount,
				'quantity', p.stock,
				'images', (
					SELECT json_agg (image_url)
					FROM product_image
					WHERE product_id = p.product_id
				)
			) AS product,
			c.quantity AS quantity
		FROM carts c
		JOIN product p ON c.product_id = p.product_id
		WHERE c.user_id = @userId
		LIMIT @limit OFFSET @offset
	`

	err := dbInstance.Raw(query, map[string]interface{}{
		"userId": userId,
		"limit":  paging.Limit,
		"offset": paging.Offset,
	}).Scan(&carts).Error

	if err != nil {
		return nil, err
	}

	var count int64
	err = dbInstance.Table("carts").Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return nil, err
	}
	paging.SetTotalPages(count)

	return carts, nil
}

func AddToCart(dbInstance *gorm.DB, userId int, productId int) error {
	query := `
		INSERT INTO carts (user_id, product_id)
		VALUES (@userId, @productId)
		ON CONFLICT (user_id, product_id) DO UPDATE
		SET quantity = carts.quantity + 1
	`
	err := dbInstance.Exec(query, map[string]interface{}{
		"userId":    userId,
		"productId": productId,
	}).Error

	if err != nil {
		return err
	}
	return nil
}
