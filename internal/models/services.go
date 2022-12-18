package models

import (
	"GoStorageService/internal/database"
	"database/sql"
	"fmt"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

// Использовать в проверке параметров запроса
var (
	dats map[string][]string = map[string][]string{
		"sbazar": {
			"electronics",
		},
	}
)

// Общая структура для обработки данных (Как для принятия так и для отдачи)
type Data struct {
	User struct {
		Name        string `json:"user_name"`
		DateRegistr string `json:"user_date_reg"`
		PhoneNumber string `json:"user_phone"`
	} `json:"user_data"`
	Products struct {
		ProdName    string `json:"product_name"`
		PhotoUrl    string `json:"photo_url"`
		Price       string `json:"price"`
		Description string `json:"description"`
		Url         string `json:"url"`
	} `json:"product_data"`
}

// Добавление данных
func AddingData(payload []byte, market string, category string) (bool, string) {
	var data []Data

	if status := CheckUrlQuery(market, category); !status {
		logrus.Warn("The parameters from the service were not correct")
		return false, "Incorrect url query"
	}

	if err := gojson.Unmarshal(payload, &data); err != nil {
		logrus.Errorf("Err unmarshal data - %s", err)
		return false, "Incorrect content data"
	}

	db, err := database.ConnectDataBase()
	if err != nil {
		return false, "Storage is not connected"
	}
	defer db.Close()

	RemoveData(db, data, market, category)

	if status, massage := AddData(db, data, market, category); !status {
		return status, massage
	}

	return true, "Success Adding"
}

// Реализованна для того чтобы пока нет реализации по обновлению данных в базе не получать дублирование результатов
func RemoveData(db *sql.DB, data []Data, market string, category string) {
	query := fmt.Sprintf("DELETE FROM %s_%s", market, category)

	if _, err := db.Exec(query); err != nil {
		logrus.Infof("Err remove data in %s_%s - %s", market, category, err)
	}
}

// Добавление новых данных
func AddData(db *sql.DB, data []Data, market string, category string) (bool, string) {
	query := fmt.Sprintf(
		"INSERT INTO %s_%s (user_name, registration_date, phone_number, product_name, product_photo_url, product_price, product_description, product_url) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", market, category)
	for _, val := range data {
		// Перед добавлением реализовать обновление уже имеющихся данных
		if _, err := db.Exec(query, val.User.Name, val.User.DateRegistr, val.User.PhoneNumber, val.Products.ProdName, val.Products.PhotoUrl, val.Products.Price, val.Products.Description, val.Products.ProdName); err != nil {
			logrus.Errorf("Err addings data to table - %s", err)
			return false, "Err add data"
		}
	}
	logrus.Infof("The %s_%s table has been updated with the following data", market, category)
	return true, ""
}

// Проверка параметров URL
func CheckUrlQuery(market string, category string) bool {
	if market == "" || category == "" {
		return false
	}

	if len(dats[market]) == 0 {
		return false
	}

	var status bool
	for _, val := range dats[market] {
		if val == category {
			status = true
		} else {
			status = false
		}
	}

	return status
}
