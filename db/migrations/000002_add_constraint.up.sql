ALTER TABLE orders ADD CONSTRAINT fk_orders_users FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE;

ALTER TABLE order_details ADD CONSTRAINT fk_details_product FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE;

ALTER TABLE order_details ADD CONSTRAINT fk_details_orders FOREIGN KEY (order_id) REFERENCES orders (order_id) ON DELETE CASCADE;

ALTER TABLE reviews ADD CONSTRAINT fk_reviews_product FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE;

ALTER TABLE reviews ADD CONSTRAINT fk_reviews_user FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE;

ALTER TABLE carts ADD CONSTRAINT pk_carts PRIMARY KEY (user_id, product_id);

ALTER TABLE carts ADD CONSTRAINT fk_carts_users FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE;

ALTER TABLE product_image ADD CONSTRAINT fk_product_image_product FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE;