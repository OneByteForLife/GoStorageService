CREATE TABLE users_telegram
(
    id serial not null unique,
    username varchar(255) not null,
    pass varchar(255) not null,
    subscription boolean
);

CREATE TABLE sbazar_electronics
(
    id serial not null unique,
    sellers_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);