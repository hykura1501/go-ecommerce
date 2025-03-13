ALTER TABLE product_image
DROP CONSTRAINT fk_product_image_product;

ALTER TABLE carts
DROP CONSTRAINT fk_carts_users;

ALTER TABLE carts
DROP CONSTRAINT pk_carts;

ALTER TABLE reviews
DROP CONSTRAINT fk_reviews_product;

ALTER TABLE order_details
DROP CONSTRAINT fk_details_orders;

ALTER TABLE order_details
DROP CONSTRAINT fk_details_product;

ALTER TABLE orders
DROP CONSTRAINT fk_orders_users;