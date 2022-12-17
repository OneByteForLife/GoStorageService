package models

import (
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

	/*
		Добовление данных в базу отталкиваясь от Query параметров (СПРОЕКТИРОВАТЬ БАЗУ)
	*/

	// out, _ := gojson.Marshal(&data)
	// fmt.Println(string(out))

	return true, "Success Adding"
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
