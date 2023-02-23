package entity

// Общая структура для обработки данных
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
