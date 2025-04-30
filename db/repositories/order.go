package repositories

import (
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateOrder(dbInstance *gorm.DB, userId int, req *entity.CreateOrderRequest) error {
	// Check product stock
	tx := dbInstance.Begin()
	for _, item := range req.Details {
		var stock int
		if err := tx.Model(&entity.Product{}).
			Where("product_id = ?", item.ProductID).
			Select("stock").
			Scan(&stock).Error; err != nil {
			tx.Rollback()
			return err
		}
		if stock < item.Quantity {
			tx.Rollback()
			return errors.New(pkg.ErrorInsufficientStock)
		}
	}

	now := time.Now()
	order := entity.Order{
		UserId:    userId,
		Status:    &pkg.OrderStatusCompleted,
		OrderDate: &now,
		Total:     &req.Total,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range req.Details {
		queryInsert := `
			INSERT INTO order_details (order_id, product_id, quantity, subtotal)
			VALUES ($1, $2, $3, $4)
		`

		if err := tx.Exec(queryInsert, order.OrderId, item.ProductID, item.Quantity, item.Subtotal).Error; err != nil {
			tx.Rollback()
			return errors.New(pkg.ErrorInsertOrderDetail)
		}

		// Update product stock
		queryUpdate := `
			UPDATE product SET stock = stock - $1
			WHERE product_id = $2
		`
		if err := tx.Exec(queryUpdate, item.Quantity, item.ProductID).Error; err != nil {
			tx.Rollback()
			return errors.New(pkg.ErrorUpdateProductStock)
		}

		// Delete cart item
		queryDelete := `
			DELETE FROM carts WHERE user_id = $1 AND product_id = $2
		`

		if err := tx.Exec(queryDelete, userId, item.ProductID).Error; err != nil {
			tx.Rollback()
			return errors.New(pkg.ErrorDeleteCartItem)
		}
	}

	tx.Commit()
	return nil
}

func GetOrderHistoryOfUser(dbInstance *gorm.DB, userId int, orderQuery *entity.OrderQuery) ([]entity.OrderHistoryReponse, error) {
	query := fmt.Sprintf(`
		SELECT 
			o.order_id,
			o.total,
			o.status,
			o.order_date,
			json_agg(json_build_object(
				'id', od.order_detail_id,
				'product', json_build_object(
					'id', p.product_id,
					'name', p.product_name,
					'price', p.price,
					'images',  (
						SELECT json_agg(image_url)
						FROM product_image
						WHERE product_id = p.product_id
					),
					'category', json_build_object(
						'category_id', c.category_id,
						'category_name', c.category_name
					),
					'manufacturer', json_build_object(
						'manufacturer_id', m.manufacturer_id,
						'manufacturer_name', m.manufacturer_name
					)
				),
				'quantity', od.quantity,
				'subtotal', od.subtotal
			)) AS details
		FROM orders o JOIN order_details od ON o.order_id = od.order_id
		JOIN product p ON od.product_id = p.product_id
		JOIN category c ON p.category_id = c.category_id
		JOIN manufacturer m ON p.manufacturer_id = m.manufacturer_id
		WHERE o.user_id = @userId
			AND (@status = '' OR o.status = @status)
			AND (@date = '' OR o.order_date::DATE = TO_DATE(@date, 'YYYY-MM-DD'))
		GROUP BY o.order_id, o.total, o.status, o.order_date
		ORDER BY %s
		LIMIT @limit OFFSET @offset
	`, orderQuery.OrderBy)
	var orderHistoryResponse []entity.OrderHistoryReponse
	params := map[string]interface{}{
		"userId": userId,
		"status": orderQuery.Status,
		"date":   orderQuery.Date,
		"limit":  orderQuery.Limit,
		"offset": orderQuery.Offset,
	}

	err := dbInstance.Raw(query, params).Scan(&orderHistoryResponse).Error
	if err != nil {
		return nil, err
	}

	var count int64
	countQuery := `
		SELECT 
			COUNT(DISTINCT o.order_id)
		FROM orders o
		WHERE o.user_id = @userId
			AND (@status = '' OR o.status = @status)
			AND (@date = '' OR o.order_date::DATE = TO_DATE(@date, 'YYYY-MM-DD'))
	`

	err = dbInstance.Raw(countQuery, params).Scan(&count).Error
	if err != nil {
		return nil, err
	}
	orderQuery.SetTotalPages(count)

	return orderHistoryResponse, nil
}

func GetOrders(dbInstance *gorm.DB, orderQuery *entity.OrderQuery) ([]entity.OrderHistoryReponse, error) {
	query := fmt.Sprintf(`
		SELECT 
			o.order_id,
			o.total,
			o.status,
			o.order_date,
			json_agg(json_build_object(
				'id', od.order_detail_id,
				'product', json_build_object(
					'id', p.product_id,
					'name', p.product_name,
					'price', p.price,
					'images',  (
						SELECT json_agg(image_url)
						FROM product_image
						WHERE product_id = p.product_id
					),
					'category', json_build_object(
						'category_id', c.category_id,
						'category_name', c.category_name
					),
					'manufacturer', json_build_object(
						'manufacturer_id', m.manufacturer_id,
						'manufacturer_name', m.manufacturer_name
					)
				),
				'quantity', od.quantity,
				'subtotal', od.subtotal
			)) AS details,
			json_build_object(
				'id', u.user_id,
				'fullname', u.fullname,
				'username', u.username,
				'avatar', u.avatar,
				'phone', u.phone,
				'address', u.address
			) as user
		FROM orders o 
		JOIN users u ON o.user_id = u.user_id
		JOIN order_details od ON o.order_id = od.order_id
		JOIN product p ON od.product_id = p.product_id
		JOIN category c ON p.category_id = c.category_id
		JOIN manufacturer m ON p.manufacturer_id = m.manufacturer_id
		WHERE (@status = '' OR o.status = @status)
		GROUP BY o.order_id, 
					o.total, 
					o.status, 
					o.order_date, 
					u.user_id, 
					u.fullname, 
					u.username, 
					u.avatar, 
					u.phone, 
					u.address
		ORDER BY %s
		LIMIT @limit OFFSET @offset
	`, orderQuery.OrderBy)
	var orderHistoryResponse []entity.OrderHistoryReponse
	params := map[string]interface{}{
		"status": orderQuery.Status,
		"limit":  orderQuery.Limit,
		"offset": orderQuery.Offset,
	}
	err := dbInstance.Raw(query, params).Scan(&orderHistoryResponse).Error
	if err != nil {
		return nil, err
	}

	var count int64
	countQuery := `
		SELECT
			COUNT(DISTINCT o.order_id)
		FROM orders o
		WHERE (@status = '' OR o.status = @status)
	`
	err = dbInstance.Raw(countQuery, params).Scan(&count).Error
	if err != nil {
		return nil, err
	}
	orderQuery.SetTotalPages(count)
	return orderHistoryResponse, nil
}

func GetRevenueStatistics(dbInstance *gorm.DB) ([]entity.RevenueStatisticsResponse, error) {
	query := `
		SELECT
			EXTRACT(MONTH FROM order_date)::INTEGER AS month,
			EXTRACT(YEAR FROM order_date)::INTEGER AS year,
			SUM(total)::REAL AS revenue
		FROM orders
		GROUP BY month, year
		LIMIT 12
	`
	var revenueStatistics []entity.RevenueStatisticsResponse
	err := dbInstance.Raw(query).Scan(&revenueStatistics).Error
	if err != nil {
		return nil, err
	}
	return revenueStatistics, nil
}

func GetBestSellersStatistics(dbInstance *gorm.DB) ([]entity.BestSellersStatisticsResponse, error) {
	query := `
		SELECT 
			p.product_id as id, 
			p.product_name as name,  
			SUM(d.quantity)::INTEGER as quantity
		FROM order_details d
		JOIN product p ON d.product_id = p.product_id
		GROUP BY p.product_id, p.product_name
		ORDER BY quantity DESC 
		LIMIT 10
	`
	var bestSellersStatistics []entity.BestSellersStatisticsResponse
	err := dbInstance.Raw(query).Scan(&bestSellersStatistics).Error
	if err != nil {
		return nil, err
	}

	return bestSellersStatistics, nil
}

func GetTopCustomersStatistics(dbInstance *gorm.DB) ([]entity.TopCustomersStatisticsResponse, error) {
	query := `
		SELECT 
			u.user_id as id, 
			u.fullname as fullname,  
			SUM(o.total)::INTEGER as total
		FROM orders o
		JOIN users u ON o.user_id = u.user_id
		GROUP BY u.user_id, u.fullname
		ORDER BY total DESC 
		LIMIT 10
	`
	var topCustomersStatistics []entity.TopCustomersStatisticsResponse
	err := dbInstance.Raw(query).Scan(&topCustomersStatistics).Error
	if err != nil {
		return nil, err
	}
	return topCustomersStatistics, nil
}

func GetOrderDetail(dbInstance *gorm.DB, userId int, orderId int, permission int) (entity.OrderHistoryReponse, error) {
	query := `
		SELECT 
			o.order_id,
			o.total,
			o.status,
			o.order_date,
			json_agg(json_build_object(
				'id', od.order_detail_id,
				'product', json_build_object(
					'id', p.product_id,
					'name', p.product_name,
					'price', p.price,
					'images', (
						SELECT json_agg(image_url)
						FROM product_image
						WHERE product_id = p.product_id
					),
					'category', json_build_object(
						'category_id', c.category_id,
						'category_name', c.category_name
					),
					'manufacturer', json_build_object(
						'manufacturer_id', m.manufacturer_id,
						'manufacturer_name', m.manufacturer_name
					)
				),
				'quantity', od.quantity,
				'subtotal', od.subtotal
			)) AS details,
			json_build_object(
				'id', u.user_id,
				'fullname', u.fullname,
				'username', u.username,
				'avatar', u.avatar,
				'phone', u.phone,
				'address', u.address
			) as user
		FROM orders o
		JOIN users u ON o.user_id = u.user_id
		JOIN order_details od ON o.order_id = od.order_id
		JOIN product p ON od.product_id = p.product_id
		JOIN category c ON p.category_id = c.category_id
		JOIN manufacturer m ON p.manufacturer_id = m.manufacturer_id
		WHERE o.order_id = $1 AND (o.user_id = $2 OR $3 = 1)
		GROUP BY o.order_id, 
				o.total, 
				o.status, 
				o.order_date, 
				u.user_id, 
				u.fullname, 
				u.username, 
				u.avatar, 
				u.phone, 
				u.address;
	`
	var order entity.OrderHistoryReponse
	err := dbInstance.Raw(query, orderId, userId, permission).Scan(&order).Error
	if err != nil {
		return entity.OrderHistoryReponse{}, nil
	}

	return order, nil
}
