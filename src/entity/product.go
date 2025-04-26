package entity

type Product struct {
	ProductId      int     `json:"product_id"`
	ProductName    string  `json:"product_name"`
	Price          float32 `json:"price"`
	Stock          int     `json:"stock"`
	Description    string  `json:"description"`
	CategoryId     int     `json:"category_id"`
	ManufacturerId string  `json:"manufacturer_id"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Discount       float32 `json:"discount"`
	Type           string  `json:"type"`
	Tag            string  `json:"tag"`
}

func (p *Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	Paging
	CategoryId int    `json:"category_id" query:"category_id"`
	Tag        string `json:"tag" query:"tag"`
	Search     string `json:"search" query:"search"`
	PriceMin   int    `json:"price_min" query:"price_min"`
	PriceMax   int    `json:"price_max" query:"price_max"`
	Sort       string `json:"sort" query:"sort"`
	OrderBy    string
}

var SortProductsOptions = map[string]string{
	"product_id_asc":    "p.product_id asc",
	"product_id_desc":   "p.product_id desc",
	"price_asc":         "p.price asc",
	"price_desc":        "p.price desc",
	"created_at_asc":    "p.created_at asc",
	"created_at_desc":   "p.created_at desc",
	"product_name_asc":  "p.product_name asc",
	"product_name_desc": "p.product_name desc",
}
