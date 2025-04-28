package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"mime/multipart"
)

type Category struct {
	CategoryId      int        `json:"category_id"`
	CategoryName    string     `json:"category_name"`
	Thumbnail       string     `json:"thumbnail"`
	Description     string     `json:"description"`
	SuperCategoryId int        `json:"super_category_id"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	Children        []Category `json:"children,omitempty" gorm:"-"`
}

func (c *Category) TableName() string {
	return "category"
}

type CategoryRequest struct {
	CategoryName    string                `form:"category_name"`
	Thumbnail       *multipart.FileHeader `form:"thumbnail" gorm:"-"`
	ThumbnailUrl    *string               `form:"-" gorm:"column:thumbnail"`
	Description     *string               `form:"description"`
	SuperCategoryId *int                  `form:"super_category_id"`
}

func (c *CategoryRequest) TableName() string {
	return "category"
}

// implement Scanner
func (c *Category) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal Category value: %v", value)
	}
	return json.Unmarshal(bytes, c)
}

// implement Valuer
func (c Category) Value() (driver.Value, error) {
	return json.Marshal(c)
}
