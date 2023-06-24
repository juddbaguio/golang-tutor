-- auto-generated definition
create table if not exists orders
(
    id          bigserial
        primary key,
    amount      numeric,
    customer_id integer
        constraint cascade_delete
            references users
            on delete cascade
);

-- auto-generated definition
create table if not exists users
(
    id         bigserial
        primary key,
    first_name varchar,
    last_name  varchar,
    username   varchar,
    password   varchar
);