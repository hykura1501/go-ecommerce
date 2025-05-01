package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Manufacturer struct {
	ManufacturerId   int    `json:"manufacturer_id"`
	ManufacturerName string `json:"manufacturer_name"`
}

func (m Manufacturer) TableName() string {
	return "manufacturer"
}

func (m *Manufacturer) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal Manufacturer value: %v", value)
	}
	return json.Unmarshal(bytes, m)
}

func (m Manufacturer) Value() (driver.Value, error) {
	return json.Marshal(m)
}
