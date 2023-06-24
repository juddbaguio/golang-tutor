select o.id as order_id, u.id as user_id, first_name, last_name, amount from users u
INNER JOIN orders o ON o.customer_id = u.id;

-- prioritizes table A have column values
select o.id as order_id, u.id as user_id, first_name, last_name, amount from users u
LEFT JOIN orders o ON o.customer_id = u.id;

-- prioritizes table B have column values
select o.id as order_id, u.id as user_id, first_name, last_name, amount from users u
RIGHT JOIN orders o ON o.customer_id = u.id;