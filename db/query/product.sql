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
    jsonb_agg(pi.image_url) AS images
FROM product p 
LEFT JOIN product_image pi ON pi.product_id = p.product_id
LEFT JOIN category c ON c.category_id = p.category_id
LEFT JOIN manufacturer m ON m.manufacturer_id = p.manufacturer_id
GROUP BY 
    p.product_id, p.product_name, p.price, p.stock, p.description, p.discount, 
    c.category_id, c.name, 
    m.manufacturer_id, m.manufacturer_name
ORDER BY p.product_id
LIMIT $1 OFFSET $2;
-- name: CountProducts :one
SELECT COUNT(*) AS total_products FROM product;

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

