package validate

import "errors"

// Использовать в проверке параметров запроса
var (
	markets map[string][]string = map[string][]string{
		"sbazar": {
			"electronics",
			"sport",
			"babymom",
			"hobby",
			"clothing",
		},
		"wallapop": {
			"electronics",
			"cars",
			"motos",
			"home",
			"phototv",
		},
		"jofogas": {
			"electronics",
			"sport",
			"babymom",
			"hobby",
			"clothing",
		},
	}
)

// Проверка параметров URL
func CheckQueryMarket(market string, category string) (error, bool) {
	if market == "" || category == "" {
		return errors.New("error query cannot be empty"), false
	}

	if len(markets[market]) == 0 {
		return errors.New("error no existing value"), false
	}

	var status bool = false
	for _, val := range markets[market] {
		if val == category {
			status = true
			return nil, status
		}
	}

	return errors.New("error no existing value"), status
}
