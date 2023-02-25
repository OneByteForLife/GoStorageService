package validate

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
func CheckQueryMarket(market string, category string) bool {
	if market == "" || category == "" {
		return false
	}

	if len(markets[market]) == 0 {
		return false
	}

	var status bool = false
	for _, val := range markets[market] {
		if val == category {
			status = true
			return status
		}
	}

	return status
}
