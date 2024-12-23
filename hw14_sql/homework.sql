CREATE TABLE Users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE UNIQUE INDEX idx_users_email ON Users (email);

CREATE TABLE Orders
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    order_date DATE NOT NULL,
    total_amount INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
);

CREATE INDEX idx_orders_user_id ON Orders (user_id);
CREATE INDEX idx_orders_order_date ON Orders (order_date);

CREATE TABLE Products
(
    id SERIAL PRIMARY KEY,
        name VARCHAR NOT NULL,
price INTEGER NOT NULL
);

CREATE INDEX idx_products_price ON Products (price);

CREATE TABLE OrderProducts
(
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    FOREIGN KEY (order_id) REFERENCES Orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products (id) ON DELETE CASCADE
);

CREATE INDEX idx_order_products_order_id ON OrderProducts (order_id);
CREATE INDEX idx_order_products_product_id ON OrderProducts (product_id);

-- Обновить информацию о пользователе
UPDATE Users
SET name = 'Dmitriy', email = 'kavo_sobaka@gmail.com', password = '123456'
WHERE id = 1;

-- Удалить пользователя
DELETE FROM Users
WHERE id = 1;

-- Добавить нового пользователя
INSERT INTO Users (name, email, password)
VALUES ('Gnom', 'dro4ila@gmail.com', '69');

-- Обновить информацию о продукте
UPDATE Products
SET name = 'M4A1', price = 100
WHERE id = 1;

-- Удалить продукт
DELETE FROM Products
WHERE id = 1;

-- Добавить новый продукт
INSERT INTO Products (name, price)
VALUES ('SSP5', 200);

-- Сохранение (Добавление заказа)
INSERT INTO Orders (user_id, order_date, total_amount)
VALUES (1, '2024-12-25', 500);

-- Удаление заказа
DELETE FROM Orders
WHERE id = 1;

-- Выборка пользователей
SELECT * FROM Users;

-- Выборка товаров
SELECT * FROM Products;

-- Выборка заказов по пользователям
SELECT Orders.id AS order_id,
       Orders.order_date,
       Orders.total_amount,
       Users.name AS user_name,
       Users.email
FROM Orders
         JOIN Users ON Orders.user_id = Users.id
WHERE Users.id = 1;

-- Выборка статистики по пользователю
SELECT Users.id AS user_id,
       Users.name AS user_name,
       SUM(Orders.total_amount) AS total_order_sum,
       AVG(Products.price) AS avg_product_price
FROM Users
         JOIN Orders ON Users.id = Orders.user_id
         JOIN OrderProducts ON Orders.id = OrderProducts.order_id
         JOIN Products ON OrderProducts.product_id = Products.id
WHERE Users.id = 1
GROUP BY Users.id, Users.name;
