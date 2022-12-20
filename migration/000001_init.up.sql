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
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE sbazar_clothing
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE sbazar_hobby
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE sbazar_sport
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE sbazar_babymom
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE wallapop_electronics
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE wallapop_cars
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE wallapop_phototv
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE wallapop_home
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);

CREATE TABLE wallapop_motos
(
    id serial not null unique,
    user_name varchar(255),
    registration_date varchar(255),
    phone_number varchar(25),
    product_name varchar(255),
    product_photo_url varchar(300),
    product_price varchar(80),
    product_description varchar,
    product_url varchar
);




