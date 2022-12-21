
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

CREATE TABLE jofogas_electronics
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

CREATE TABLE jofogas_clothing
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

CREATE TABLE jofogas_hobby
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

CREATE TABLE jofogas_sport
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

CREATE TABLE jofogas_babymom
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


