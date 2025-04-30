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

type RevenueStatisticsResponse struct {
	Month   int     `json:"month"`
	Year    int     `json:"year"`
	Revenue float32 `json:"revenue"`
}

type BestSellersStatisticsResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type TopCustomersStatisticsResponse struct {
	Id       int     `json:"id"`
	Fullname string  `json:"fullname"`
	Total    float32 `json:"total"`
}
