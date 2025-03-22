-- Insert users (10 records)
INSERT INTO users (username, password, permission, login_provider, provider_id, account_id, fullname, phone, address)
VALUES 
    ('alice', 'hashed_pw_1', 1, NULL, NULL, 1, 'Alice Johnson', '0123456789', '123 Main St'),
    ('bob', 'hashed_pw_2', 1, NULL, NULL, 2, 'Bob Smith', '0987654321', '456 Elm St'),
    ('charlie', 'hashed_pw_3', 1, NULL, NULL, 3, 'Charlie Brown', '0111222333', '789 Pine St'),
    ('david', 'hashed_pw_4', 1, NULL, NULL, 4, 'David Williams', '0222333444', '101 Oak St'),
    ('eva', 'hashed_pw_5', 1, NULL, NULL, 5, 'Eva Miller', '0333444555', '202 Maple St'),
    ('frank', 'hashed_pw_6', 1, NULL, NULL, 6, 'Frank Harris', '0444555666', '303 Birch St'),
    ('grace', 'hashed_pw_7', 1, NULL, NULL, 7, 'Grace Lee', '0555666777', '404 Cedar St'),
    ('hannah', 'hashed_pw_8', 1, NULL, NULL, 8, 'Hannah Davis', '0666777888', '505 Walnut St'),
    ('isaac', 'hashed_pw_9', 1, NULL, NULL, 9, 'Isaac White', '0777888999', '606 Elm St'),
    ('julia', 'hashed_pw_10', 1, NULL, NULL, 10, 'Julia Adams', '0888999000', '707 Chestnut St');

-- Insert accounts (10 records)
INSERT INTO accounts (balance)
VALUES (1000.00), (1500.00), (1200.50), (2000.75), (1800.30), 
       (500.00), (2500.20), (1750.80), (900.60), (3000.00);

-- Insert manufacturers (10 records)
INSERT INTO manufacturer (manufacturer_name)
VALUES 
    ('Samsung'), ('Apple'), ('Sony'), ('LG'), ('Dell'),
    ('HP'), ('Microsoft'), ('Lenovo'), ('Asus'), ('Acer');

-- Insert categories (10 records)
INSERT INTO category (name, thumbnail, description)
VALUES 
    ('Electronics', NULL, 'Electronic devices'),
    ('Mobile Phones', NULL, 'Smartphones and accessories'),
    ('Laptops', NULL, 'Personal and gaming laptops'),
    ('Tablets', NULL, 'Android and iOS tablets'),
    ('Wearables', NULL, 'Smartwatches and fitness bands'),
    ('Audio', NULL, 'Headphones and speakers'),
    ('Cameras', NULL, 'DSLR and mirrorless cameras'),
    ('Gaming', NULL, 'Gaming consoles and accessories'),
    ('Accessories', NULL, 'Cables, chargers, and cases'),
    ('Smart Home', NULL, 'Smart devices for home automation');

-- Insert products (20 records)
INSERT INTO product (product_name, price, stock, description, category_id, manufacturer_id, discount, type, tag)
VALUES 
    ('iPhone 15', 999.99, 50, 'Latest iPhone model', 2, 2, 5, 'Smartphone', 'new'),
    ('Samsung Galaxy S24', 899.99, 40, 'Samsung flagship', 2, 1, 10, 'Smartphone', 'sale'),
    ('Sony WH-1000XM5', 399.99, 100, 'Noise-canceling headphones', 6, 3, 15, 'Headphones', 'featured'),
    ('MacBook Pro 14', 1999.99, 30, 'Apple M3 chip', 3, 2, 0, 'Laptop', 'new'),
    ('Dell XPS 15', 1499.99, 25, 'High-performance laptop', 3, 5, 5, 'Laptop', 'sale'),
    ('Apple iPad Air', 599.99, 80, 'Latest iPad Air', 4, 2, 10, 'Tablet', 'featured'),
    ('Samsung Galaxy Tab S9', 699.99, 70, 'Samsung flagship tablet', 4, 1, 10, 'Tablet', 'new'),
    ('Garmin Forerunner 955', 499.99, 60, 'Premium running smartwatch', 5, 7, 20, 'Smartwatch', 'sale'),
    ('Sony Alpha A7 IV', 2499.99, 20, 'Full-frame mirrorless camera', 7, 3, 5, 'Camera', 'featured'),
    ('Bose SoundLink', 149.99, 90, 'Portable Bluetooth speaker', 6, 10, 10, 'Speaker', 'new'),
    ('Asus ROG Phone 7', 799.99, 35, 'Gaming smartphone', 2, 9, 15, 'Smartphone', 'sale'),
    ('Lenovo ThinkPad X1', 1299.99, 40, 'Business laptop', 3, 8, 5, 'Laptop', 'featured'),
    ('Microsoft Surface Pro 9', 1199.99, 30, '2-in-1 laptop', 4, 6, 10, 'Tablet', 'new'),
    ('Google Pixel Watch 2', 349.99, 75, 'Wearable from Google', 5, 7, 15, 'Smartwatch', 'sale'),
    ('Nintendo Switch OLED', 349.99, 50, 'Gaming console', 8, 7, 0, 'Gaming Console', 'featured'),
    ('Samsung QLED 65"', 1299.99, 20, '65-inch 4K QLED TV', 10, 1, 5, 'TV', 'new'),
    ('HP Spectre x360', 1399.99, 35, 'Convertible laptop', 3, 6, 10, 'Laptop', 'sale'),
    ('Canon EOS R5', 3799.99, 15, 'Professional mirrorless camera', 7, 4, 5, 'Camera', 'featured'),
    ('Logitech MX Master 3', 99.99, 100, 'Wireless mouse', 9, 8, 10, 'Mouse', 'new'),
    ('Razer BlackWidow V4', 199.99, 80, 'Mechanical gaming keyboard', 9, 9, 5, 'Keyboard', 'sale');



INSERT INTO product_image (product_id, image_url) VALUES
(1, 'https://example.com/images/product1_1.jpg'),
(1, 'https://example.com/images/product1_2.jpg'),
(2, 'https://example.com/images/product2_1.jpg'),
(2, 'https://example.com/images/product2_2.jpg'),
(2, 'https://example.com/images/product2_3.jpg'),
(3, 'https://example.com/images/product3_1.jpg'),
(3, 'https://example.com/images/product3_2.jpg'),
(3, 'https://example.com/images/product3_3.jpg'),
(4, 'https://example.com/images/product4_1.jpg'),
(4, 'https://example.com/images/product4_2.jpg'),
(5, 'https://example.com/images/product5_1.jpg'),
(5, 'https://example.com/images/product5_2.jpg'),
(5, 'https://example.com/images/product5_3.jpg'),
(5, 'https://example.com/images/product5_4.jpg'),
(6, 'https://example.com/images/product6_1.jpg'),
(6, 'https://example.com/images/product6_2.jpg'),
(7, 'https://example.com/images/product7_1.jpg'),
(7, 'https://example.com/images/product7_2.jpg'),
(7, 'https://example.com/images/product7_3.jpg'),
(8, 'https://example.com/images/product8_1.jpg'),
(8, 'https://example.com/images/product8_2.jpg'),
(9, 'https://example.com/images/product9_1.jpg'),
(9, 'https://example.com/images/product9_2.jpg'),
(9, 'https://example.com/images/product9_3.jpg'),
(9, 'https://example.com/images/product9_4.jpg'),
(10, 'https://example.com/images/product10_1.jpg'),
(10, 'https://example.com/images/product10_2.jpg'),
(10, 'https://example.com/images/product10_3.jpg'),
(11, 'https://example.com/images/product11_1.jpg'),
(11, 'https://example.com/images/product11_2.jpg'),
(12, 'https://example.com/images/product12_1.jpg'),
(12, 'https://example.com/images/product12_2.jpg'),
(12, 'https://example.com/images/product12_3.jpg'),
(12, 'https://example.com/images/product12_4.jpg'),
(13, 'https://example.com/images/product13_1.jpg'),
(13, 'https://example.com/images/product13_2.jpg'),
(14, 'https://example.com/images/product14_1.jpg'),
(14, 'https://example.com/images/product14_2.jpg'),
(14, 'https://example.com/images/product14_3.jpg'),
(15, 'https://example.com/images/product15_1.jpg'),
(15, 'https://example.com/images/product15_2.jpg'),
(16, 'https://example.com/images/product16_1.jpg'),
(16, 'https://example.com/images/product16_2.jpg'),
(16, 'https://example.com/images/product16_3.jpg'),
(17, 'https://example.com/images/product17_1.jpg'),
(17, 'https://example.com/images/product17_2.jpg'),
(18, 'https://example.com/images/product18_1.jpg'),
(18, 'https://example.com/images/product18_2.jpg'),
(18, 'https://example.com/images/product18_3.jpg'),
(19, 'https://example.com/images/product19_1.jpg'),
(19, 'https://example.com/images/product19_2.jpg'),
(20, 'https://example.com/images/product20_1.jpg'),
(20, 'https://example.com/images/product20_2.jpg'),
(20, 'https://example.com/images/product20_3.jpg'),
(20, 'https://example.com/images/product20_4.jpg'),
(20, 'https://example.com/images/product20_5.jpg');


-- Insert orders (10 records)
INSERT INTO orders (total, status, user_id)
VALUES 
    (1999.98, 'completed', 1), (899.99, 'pending', 2), (499.99, 'shipped', 3),
    (2499.99, 'delivered', 4), (149.99, 'completed', 5), (349.99, 'pending', 6),
    (1299.99, 'completed', 7), (1199.99, 'shipped', 8), (799.99, 'delivered', 9), (999.99, 'pending', 10);

-- Insert order details (10 records)
INSERT INTO order_details (product_id, quantity, subtotal, order_id)
VALUES 
    (1, 2, 1999.98, 1), (2, 1, 899.99, 2), (6, 1, 499.99, 3),
    (9, 1, 2499.99, 4), (10, 1, 149.99, 5), (14, 1, 349.99, 6),
    (15, 1, 1299.99, 7), (16, 1, 1199.99, 8), (17, 1, 799.99, 9), (18, 1, 999.99, 10);
