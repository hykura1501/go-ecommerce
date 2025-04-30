package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ProductCart struct {
	Id       int          `json:"id,omitempty"`
	Name     string       `json:"name,omitempty"`
	Price    float32      `json:"price,omitempty"`
	Quantity int          `json:"quantity,omitempty"`
	Images   *StringArray `json:"images,omitempty"`
	Discount *float32     `json:"discount,omitempty"`
}

type Cart struct {
	Product  ProductCart `json:"product" gorm:"column:product"`
	Quantity int         `json:"quantity"`
}

type AddToCartRequest struct {
	ProductId int `json:"product_id"`
}

// implement Scanner
func (c *ProductCart) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal ProductCart value: %v", value)
	}
	return json.Unmarshal(bytes, c)
}

// implement Valuer
func (c ProductCart) Value() (driver.Value, error) {
	return json.Marshal(c)
}
