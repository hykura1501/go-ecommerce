package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Custom type cho Images
type StringArray []string

// Implement Scanner
func (a *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal StringArray value: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

// Implement Valuer (nếu cần insert/update)
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}
