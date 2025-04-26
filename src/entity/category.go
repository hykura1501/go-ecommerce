package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
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
