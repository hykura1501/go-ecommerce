CREATE TABLE
    IF NOT EXISTS users (
        user_id SERIAL PRIMARY KEY,
        username VARCHAR(40) UNIQUE,
        password VARCHAR(255),
        permission INT NOT NULL,
        login_provider INT,
        provider_id VARCHAR(255),
        account_id INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        fullname VARCHAR(255),
        phone VARCHAR(20),
        address VARCHAR(255),
        avatar TEXT DEFAULT 'https://res.cloudinary.com/dnrz2djhd/image/upload/v1734689215/d9wy5b4anfqh1rksfiwk.png'
    );

CREATE TABLE
    IF NOT EXISTS reviews (
        review_id SERIAL PRIMARY KEY,
        content TEXT NOT NULL,
        rating INTEGER NOT NULL,
        posted_at TIMESTAMP DEFAULT NOW (),
        product_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS carts (
        user_id INT,
        product_id INT,
        quantity INT DEFAULT 1
    );

CREATE TABLE
    IF NOT EXISTS accounts (
        account_id SERIAL PRIMARY KEY,
        balance NUMERIC(10, 2) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS transactions (
        transaction_id SERIAL PRIMARY KEY,
        account_id INT NOT NULL,
        amount NUMERIC(10, 2) NOT NULL,
        status VARCHAR(20) DEFAULT 'pending',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        order_id INT NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS manufacturer (
        manufacturer_id SERIAL PRIMARY KEY,
        manufacturer_name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS category (
        category_id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        thumbnail VARCHAR(255),
        description TEXT,
        super_category_id INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_super_category FOREIGN KEY (super_category_id) REFERENCES category (category_id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS product (
        product_id SERIAL PRIMARY KEY,
        product_name VARCHAR(255) NOT NULL,
        price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
        stock INT NOT NULL CHECK (stock >= 0),
        description TEXT,
        category_id INT NOT NULL,
        manufacturer_id INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        discount NUMERIC(5, 2) CHECK (
            discount >= 0
            AND discount <= 100
        ) DEFAULT 0,
        type VARCHAR(100),
        tag VARCHAR(100),
        CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category (category_id) ON DELETE CASCADE,
        CONSTRAINT fk_manufacturer FOREIGN KEY (manufacturer_id) REFERENCES manufacturer (manufacturer_id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS attributes (
        attribute_id SERIAL PRIMARY KEY,
        attribute_name VARCHAR(255) NOT NULL,
        value VARCHAR(255) NOT NULL,
        product_id INT NOT NULL,
        CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE
    );
CREATE TABLE
    IF NOT EXISTS product_image (
        product_id INT NOT NULL,
        image_url TEXT NOT NULL,
        PRIMARY KEY (product_id, image_url)
    );

CREATE TABLE
    IF NOT EXISTS orders (
        order_id SERIAL PRIMARY KEY,
        total NUMERIC(10, 2),
        status VARCHAR(20) default 'pending',
        order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_id INT
    );

CREATE TABLE
    IF NOT EXISTS order_details (
        order_detail_id SERIAL PRIMARY KEY,
        product_id INT,
        quantity INT,
        subtotal NUMERIC(10, 2),
        order_id INT
    );