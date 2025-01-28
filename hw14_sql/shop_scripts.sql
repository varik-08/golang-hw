-- Напишите запросы на вставку, редактирование и удаление пользователей и продуктов.
insert into users (name, email, password)
values ('user1', 'user1@ya.ru', 'password1'),
       ('user2', 'user2@ya.ru', 'password2'),
       ('user3', 'user3@ya.ru', 'password3'),
       ('user4', 'user4@ya.ru', 'password4');

insert into products (name, price)
values ('product1', 100),
       ('product2', 200),
       ('product3', 300),
       ('product4', 400);

update users
set name = 'user0'
where email = 'user1@ya.ru';

update products
set price = 300
where name = 'product1';

delete
from users
where email = 'user2@ya.ru';
delete
from products
where price < 100;

-- Напишите запрос на сохранение и удаление заказов
insert into orders (user_id, order_date, total_amount)
values (1, '2022-01-01', 100),
       (1, '2022-01-01', 200),
       (1, '2022-01-01', 300),
       (3, '2022-01-01', 400);

delete
from orders
where user_id = 1;

-- Напишите запрос на выборку пользователей и выборку товаров
select id, password
from users
where email = 'user1@ya.ru';

select id, name
from products
where price > 100;

-- Напишите запрос на выборку заказов по пользователю
select o.id, o.order_date, p.id, p.name
from orders o
         left join order_product op on o.id = op.order_id
         left join products p on p.id = op.product_id
         left join users u on u.id = o.user_id
where u.email = 'user1@ya.ru';

-- Напишите запрос на выборку статистики по пользователю (общая сумма заказов/средняя цена товара)
select u.id, sum(o.total_amount), avg(p.price)
from users u
         left join orders o on o.user_id = u.id
         left join order_product op on o.id = op.order_id
         left join products p on p.id = op.product_id
group by u.id;

-- Создайте требуемые индексы для ускорения выборки
create index idx_users_email on users (email);
create index idx_orders_user_id_order_date on orders (user_id, order_date);
