package entity

type StatisticByCategory struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type StatisticByManufacturer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type NewUserStatistic struct {
	Month    int `json:"month"`
	Year     int `json:"year"`
	Quantity int `json:"quantity"`
}
