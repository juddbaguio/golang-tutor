ALTER TABLE orders
add constraint fk_order_user_id foreign key (customer_id) references users(id);

ALTER TABLE orders
add constraint cascade_delete foreign key (customer_id) references users(id) on DELETE CASCADE;