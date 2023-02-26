package usecase

import (
	"GoStorageService/internal/entity"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

// tables
const (
	ads       = "ads"
	ads_users = "ads_users"
	markets   = "markets"
)

// columns
const (
	ad_user_id                = "ad_user_id"
	ad_user_name              = "ad_user_name"
	ad_user_registration_date = "ad_user_registration_date"
	ad_phone_number           = "ad_phone_number"
	ad_id                     = "ad_id"
	ad_title                  = "ad_title"
	ad_photo_url              = "ad_photo_url"
	ad_price                  = "ad_price"
	ad_description            = "ad_description"
	ad_url                    = "ad_url"
	market_name               = "market_name"
	market_category           = "market_category"
	fk_ad_user_id             = "fk_ad_user_id"
	fk_ad_id                  = "fk_ad_id"
)

type MarketStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &MarketStorage{db: db}
}

func (s *MarketStorage) Add(data []entity.Data, name string, category string) error {
	if err := s.clear(); err != nil {
		return err
	}

	var Uid, Aid string
	for _, val := range data {
		tx, err := s.db.Begin()
		if err != nil {
			return err
		}

		// Insert data ads_users
		ds := goqu.Insert(ads_users).
			Cols(ad_user_name, ad_user_registration_date, ad_phone_number).Vals(
			goqu.Vals{
				&val.User.Name, &val.User.DateRegistr, &val.User.PhoneNumber,
			}).
			Returning(ad_user_id)

		q, _, _ := ds.ToSQL()

		rows, err := tx.Query(q)
		if err != nil {
			return err
		}

		for rows.Next() {
			if err := rows.Scan(&Uid); err != nil {
				return err
			}
		}

		// Insert data ads
		ds = goqu.Insert(ads).
			Cols(ad_title, ad_photo_url, ad_price, ad_description, ad_url).
			Vals(goqu.Vals{
				&val.Products.ProdName, &val.Products.PhotoUrl, &val.Products.Price, &val.Products.Description, &val.Products.Url,
			}).
			Returning(ad_id)

		q, _, _ = ds.ToSQL()

		rows, err = tx.Query(q)
		if err != nil {
			return err
		}

		for rows.Next() {
			if err := rows.Scan(&Aid); err != nil {
				return err
			}
		}

		// Insert data markets
		ds = goqu.Insert(markets).
			Cols(market_name, market_category, fk_ad_user_id, fk_ad_id).
			Vals(goqu.Vals{
				&name, &category, &Uid, &Aid,
			})

		q, _, _ = ds.ToSQL()

		_, err = tx.Exec(q)
		if err != nil {
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MarketStorage) clear() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	if err := deltable(tx, markets); err != nil {
		return err
	}
	if err := deltable(tx, ads_users); err != nil {
		return err
	}
	if err := deltable(tx, ads); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func deltable(tx *sql.Tx, table string) error {
	ds := goqu.Delete(table)
	q, _, _ := ds.ToSQL()

	_, err := tx.Exec(q)
	if err != nil {
		return err
	}

	return nil
}
