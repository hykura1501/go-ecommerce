package repositories

import (
	"BE_Ecommerce/src/entity"
	"fmt"

	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB, query *entity.ProductQuery) ([]entity.Product, error) {
	products := []entity.Product{}
	querySQL := fmt.Sprintf(`
		SELECT 
			p.product_id, 
			p.product_name, 
			p.price, 
			p.stock, 
			p.description, 
			p.discount, 
			p.tag,
			jsonb_build_object('category_id', c.category_id, 'category_name', c.name) AS category,
			jsonb_build_object('manufacturer_id', m.manufacturer_id, 'manufacturer_name', m.manufacturer_name) AS manufacturer,
			COALESCE(jsonb_agg(pi.image_url) FILTER (WHERE pi.image_url IS NOT NULL), '[]'::jsonb) AS images
		FROM product p 
		LEFT JOIN product_image pi ON pi.product_id = p.product_id
		LEFT JOIN category c ON c.category_id = p.category_id
		LEFT JOIN manufacturer m ON m.manufacturer_id = p.manufacturer_id
		WHERE
			(p.category_id = COALESCE(NULLIF(@categoryId, 0), p.category_id))
			AND (p.tag = COALESCE(NULLIF(@tag, ''), p.tag))
			AND (p.price >= COALESCE(NULLIF(@priceMin, 0), p.price))
			AND (p.price <= COALESCE(NULLIF(@priceMax, 0), p.price))
			AND (p.product_name ILIKE COALESCE(NULLIF('%%' || @search || '%%', '%%%%'), p.product_name))
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.name, 
			m.manufacturer_id, m.manufacturer_name
		ORDER BY %s
		LIMIT @limit OFFSET @offset;
	`, query.OrderBy)

	params := map[string]interface{}{
		"categoryId": query.CategoryId,
		"priceMin":   query.PriceMin,
		"priceMax":   query.PriceMax,
		"search":     query.Search,
		"limit":      query.Limit,
		"offset":     query.Offset,
		"tag":        query.Tag,
	}

	err := db.Raw(querySQL, params).Scan(&products).Error

	if err != nil {
		return nil, err
	}

	var count int64
	countSQL := `
		SELECT COUNT(*)
		FROM product p 
		WHERE
			(p.category_id = COALESCE(NULLIF(@categoryId, 0), p.category_id))
			AND (p.tag = COALESCE(NULLIF(@tag, ''), p.tag))
			AND (p.price >= COALESCE(NULLIF(@priceMin, 0), p.price))
			AND (p.price <= COALESCE(NULLIF(@priceMax, 0), p.price))
			AND (p.product_name ILIKE COALESCE(NULLIF('%%' || @search || '%%', '%%%%'), p.product_name))
	`
	err = db.Raw(countSQL, params).Scan(&count).Error

	if err != nil {
		return nil, err
	}

	query.SetTotalPages(count)

	return products, nil
}
