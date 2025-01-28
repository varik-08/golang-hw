-- Таблица ""Пользователи"" (Users) содержит информацию о пользователях, включая их
-- уникальные идентификаторы (id),
-- имена (name),
-- электронные адреса (email)
-- пароли (password).
create table users
(
    id       serial primary key,
    name     varchar(255),
    email    varchar(255),
    password varchar(255)
);


-- Таблица ""Заказы"" (Orders) отображает информацию о заказах, включая
-- идентификаторы заказов (id),
-- идентификаторы пользователей (user_id),
-- даты заказов (order_date),
-- общую стоимость заказов (total_amount).
-- Связь между таблицами ""Пользователи"" и ""Заказы"" реализована через внешний ключ (FOREIGN KEY).
create table orders
(
    id           serial primary key,
    user_id      int references users (id),
    order_date   date,
    total_amount float
);

-- Таблица ""Товары"" (Products) содержит информацию о товарах, включая их
-- идентификаторы (id),
-- названия (name)
-- цены (price).
create table products
(
    id    serial primary key,
    name  varchar(255),
    price float
);

-- Таблица ""Заказы-Товары"" (OrderProducts) содержит информацию о отношении заказов к товарам (многие ко многим).
create table order_product
(
    order_id   int references orders (id),
    product_id int references products (id)
);