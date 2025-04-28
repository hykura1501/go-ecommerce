package entity

import "mime/multipart"

type Product struct {
	ProductId    int          `json:"product_id"`
	ProductName  string       `json:"product_name"`
	Price        float32      `json:"price"`
	Stock        int          `json:"stock"`
	Description  *string      `json:"description"`
	Category     Category     `json:"category"`
	Manufacturer Manufacturer `json:"manufacturer"`
	Images       *StringArray `json:"images"`
	CreatedAt    *string      `json:"created_at"`
	UpdatedAt    *string      `json:"updated_at"`
	Discount     *float32     `json:"discount"`
	Type         *string      `json:"type"`
	Tag          *string      `json:"tag"`
}

func (p *Product) TableName() string {
	return "product"
}

type ProductImage struct {
	ProductId int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
}

func (*ProductImage) TableName() string {
	return "product_image"
}

// name, price, description, manufacturer_id, category_id, images, stock, discount, tag
type NewProductRequest struct {
	ProductId      int                     `gorm:"column:product_id;primaryKey;autoIncrement" json:"product_id"`
	Name           string                  `form:"product_name" gorm:"column:product_name" validate:"required"`
	Price          float32                 `form:"price" gorm:"column:price" validate:"required"`
	Description    *string                 `form:"description" gorm:"column:description"`
	ManufacturerId int                     `form:"manufacturer_id" gorm:"column:manufacturer_id" validate:"required"`
	CategoryId     int                     `form:"category_id" gorm:"column:category_id" validate:"required"`
	Images         []*multipart.FileHeader `form:"images" gorm:"-"`
	ImageUrls      []string                `form:"-" gorm:"-"`
	Stock          int                     `form:"stock" gorm:"column:stock" validate:"required"`
	Discount       *float32                `form:"discount" gorm:"column:discount"`
	Tag            *string                 `form:"tag" gorm:"column:tag"`
}

func (n *NewProductRequest) TableName() string {
	return "product"
}

type UpdateProductRequest struct {
	Name           *string                 `form:"product_name" gorm:"column:product_name"`
	Price          *float32                `form:"price" gorm:"column:price"`
	Stock          *int                    `form:"stock" gorm:"column:stock"`
	Description    *string                 `form:"description" gorm:"column:description"`
	ManufacturerId *int                    `form:"manufacturer_id" gorm:"column:manufacturer_id"`
	CategoryId     *int                    `form:"category_id" gorm:"column:category_id"`
	Discount       *float32                `form:"discount" gorm:"column:discount"`
	Tag            *string                 `form:"tag" gorm:"column:tag"`
	OldImageUrls   []string                `form:"old_image_urls" gorm:"-"`
	Images         []*multipart.FileHeader `form:"images" gorm:"-"`
	ImageUrls      []string                `form:"-" gorm:"-"`
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

type SpecialProductList struct {
	NewArrival      []Product `json:"new_arrival"`
	BestSeller      []Product `json:"best_seller"`
	Featured        []Product `json:"featured"`
	HighestDiscount []Product `json:"highest_discount"`
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
