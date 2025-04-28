package repositories

import (
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"errors"
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

func GetSpecialProducts(db *gorm.DB, size int) (entity.SpecialProductList, error) {
	results := entity.SpecialProductList{}
	newArivalSQL := `
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
		WHERE p.tag = 'new'
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.name, 
			m.manufacturer_id, m.manufacturer_name
		ORDER BY p.created_at DESC
		LIMIT $1;
	`
	db.Raw(newArivalSQL, size).Scan(&results.NewArrival)

	bestSellerSQL := `
		WITH best_sellers AS (
			SELECT pr.product_id
			FROM product pr
			LEFT JOIN order_details od ON pr.product_id = od.product_id
			GROUP BY pr.product_id
			ORDER BY COUNT(od.product_id) DESC
			LIMIT $1 
		)
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
		WHERE p.product_id IN (SELECT product_id FROM best_sellers)
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.name, 
			m.manufacturer_id, m.manufacturer_name
	`
	db.Raw(bestSellerSQL, size).Scan(&results.BestSeller)

	featuredSQL := `
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
		WHERE p.tag = 'featured'
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.name, 
			m.manufacturer_id, m.manufacturer_name
		LIMIT $1;
	`
	db.Raw(featuredSQL, size).Scan(&results.Featured)

	highestDiscountSQL := `
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
		WHERE p.discount > 0
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.name, 
			m.manufacturer_id, m.manufacturer_name
		ORDER BY p.discount DESC
		LIMIT $1;
	`
	db.Raw(highestDiscountSQL, size).Scan(&results.HighestDiscount)

	return results, nil
}

func GetProductDetail(db *gorm.DB, productId int) (entity.Product, error) {
	product := entity.Product{}
	querySQL := `
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
		WHERE p.product_id = $1
		GROUP BY 
			p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
			c.category_id, c.name, 
			m.manufacturer_id, m.manufacturer_name
	`
	result := db.Raw(querySQL, productId).Scan(&product)

	if result.RowsAffected == 0 {
		return entity.Product{}, errors.New(pkg.ErrorRecordNotFound)
	}

	if result.Error != nil {
		return entity.Product{}, result.Error
	}
	return product, nil
}

func CreateProduct(db *gorm.DB, product *entity.NewProductRequest) error {
	transaction := db.Begin()
	err := transaction.Table("product").Create(product).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	productId := product.ProductId
	var imageUrls []entity.ProductImage
	for _, imageUrl := range product.ImageUrls {
		imageUrls = append(imageUrls, entity.ProductImage{
			ProductId: productId,
			ImageUrl:  imageUrl,
		})
	}
	err = transaction.Table("product_image").Create(&imageUrls).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
