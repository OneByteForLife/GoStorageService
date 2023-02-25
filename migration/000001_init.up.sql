CREATE TABLE ads_users
(
    ad_user_id bigserial,
    ad_user_name varchar(255),
    ad_user_registration_date varchar(255),
    ad_phone_number varchar(25),

    CONSTRAINT PK_ads_users_ad_user_id PRIMARY KEY (ad_user_id)
);

CREATE TABLE ads
(
    ad_id bigserial,
    ad_title varchar(255),
    ad_photo_url varchar(300),
    ad_price varchar(80),
    ad_description text,
    ad_url varchar,
    
    CONSTRAINT PK_ads_ad_id PRIMARY KEY(ad_id)
);


-- MAIN TABLE
CREATE TABLE markets
(
    market_id bigserial,
    market_name varchar(100) NOT NULL,
    market_category varchar(100) NOT NULL,
    fk_ad_user_id bigint, -- Внешний ключ для пользователя
    fk_ad_id bigint, -- Внешний ключ для товаров

    CONSTRAINT PK_markets_market_id PRIMARY KEY (market_id),
    CONSTRAINT FK_ads_user_ad_user_id FOREIGN KEY (fk_ad_user_id) REFERENCES ads_users(ad_user_id),
    CONSTRAINT FK_ads_ad_id FOREIGN KEY (fk_ad_id) REFERENCES ads(ad_id)
);

