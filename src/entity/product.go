package entity

type ProductRequest struct {
	Page       int32  `json:"page" query:"page"`
	PerPage    int32  `json:"per_page" query:"per_page"`
	CategoryId int32  `json:"category_id" query:"category_id"`
	Tag        string `json:"tag" query:"tag"`
	Search     string `json:"search" query:"search"`
	PriceMin   int32  `json:"price_min" query:"price_min"`
	PriceMax   int32  `json:"price_max" query:"price_max"`
	Sort       string `json:"sort" query:"sort"`
	SortBy     string
	SortValue  string
}

var SortProductsOptions = map[string]string{
	"product_id_asc":    "product_id asc",
	"product_id_desc":   "product_id desc",
	"price_asc":         "price asc",
	"price_desc":        "price desc",
	"created_at_asc":    "created_at asc",
	"created_at_desc":   "created_at desc",
	"product_name_asc":  "product_name asc",
	"product_name_desc": "product_name desc",
}
