package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	UserId    int        `json:"user_id"`
	OrderId   int        `json:"order_id" gorm:"primaryKey"`
	Total     *float32   `json:"total"`
	Status    *string    `json:"status"`
	OrderDate *time.Time `json:"order_date"`
}

func (o Order) TableName() string {
	return "orders"
}

type OrderDetail struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Subtotal  float32 `json:"subtotal"`
}

func (o OrderDetail) TableName() string {
	return "order_details"
}

type CreateOrderRequest struct {
	Total   float32       `json:"total"`
	Details []OrderDetail `json:"details"`
}

type OrderQuery struct {
	Paging
	Order   string `json:"order"`
	Status  string `json:"status"`
	Date    string `json:"date"`
	OrderBy string
}

type ProductOrderItem struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Price        float32      `json:"price"`
	Images       StringArray  `json:"images"`
	Category     Category     `json:"category"`
	Manufacturer Manufacturer `json:"manufacturer"`
}

func (p *ProductOrderItem) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal ProductOrderItem: %v", value)
	}
	return json.Unmarshal(bytes, p)
}

func (p ProductOrderItem) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type OrderItems struct {
	Id       int              `json:"id"`
	Product  ProductOrderItem `json:"product"`
	Quantity int              `json:"quantity"`
	Subtotal float32          `json:"subtotal"`
}

type OrderDetails []OrderItems

type OrderUser struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Fulname  string `json:"fullname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
}

type OrderHistoryReponse struct {
	OrderId   int          `json:"order_id"`
	Total     float32      `json:"total"`
	Status    string       `json:"status"`
	OrderDate *time.Time   `json:"order_date"`
	Details   OrderDetails `json:"details"`
	User      OrderUser    `json:"user"`
}

func (m *OrderDetails) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal OrderDetails value: %v", value)
	}
	return json.Unmarshal(bytes, m)
}

func (m OrderDetails) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *OrderUser) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal OrderUser value: %v", value)
	}
	return json.Unmarshal(bytes, m)
}

func (m OrderUser) Value() (driver.Value, error) {
	return json.Marshal(m)
}

var OrderSortOptions = map[string]string{
	"date_asc":   "o.order_date ASC",
	"date_desc":  "o.order_date DESC",
	"total_asc":  "o.total ASC",
	"total_desc": "o.total DESC",
	"id_asc":     "o.order_id ASC",
	"id_desc":    "o.order_id DESC",
}
