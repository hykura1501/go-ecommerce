-- name: GetAllProducts :many
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
WHERE (sqlc.narg(category_id)::integer = 0 OR p.category_id = sqlc.narg(category_id)) 
AND (sqlc.narg(tag)::text = '' OR p.tag = sqlc.narg(tag)) 
AND (sqlc.narg(price_min)::integer = 0 OR p.price >= sqlc.narg(price_min))
AND (sqlc.narg(price_max)::integer = 0 OR p.price <= sqlc.narg(price_max))
AND (sqlc.narg(search)::text = '' OR p.product_name ILIKE CONCAT('%', sqlc.narg(search), '%'))
GROUP BY 
    p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
    c.category_id, c.name, 
    m.manufacturer_id, m.manufacturer_name
ORDER BY sqlc.narg(sort_by) sqlc.narg(sort_value)
LIMIT $1 OFFSET $2;


-- name: CountProducts :one
SELECT COUNT(DISTINCT p.product_id) AS total_products 
FROM product p 
LEFT JOIN product_image pi ON pi.product_id = p.product_id
LEFT JOIN category c ON c.category_id = p.category_id
LEFT JOIN manufacturer m ON m.manufacturer_id = p.manufacturer_id
WHERE (sqlc.narg(category_id)::integer = 0 OR p.category_id = sqlc.narg(category_id)) 
AND (sqlc.narg(tag)::text = '' OR p.tag = sqlc.narg(tag)) 
AND (sqlc.narg(price_min)::integer = 0 OR p.price >= sqlc.narg(price_min))
AND (sqlc.narg(price_max)::integer = 0 OR p.price <= sqlc.narg(price_max))
AND (sqlc.narg(search)::text = '' OR p.product_name ILIKE CONCAT('%', sqlc.narg(search), '%'));

-- name: GetNewArrivalProducts :many
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
ORDER BY p.product_id
LIMIT 5;

