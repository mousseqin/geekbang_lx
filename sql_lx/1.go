package sql_lx

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type Args map[string]string

// Scan Scanner
func (args Args) Scan(value interface{}) error {
	return scan(&args, value)
}

// Value Valuer
func (args Args) Value() (driver.Value, error) {
	return value(args), nil
}

// scan for scanner helper
func scan(data interface{}, value interface{}) error {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), data)
	case string:
		return json.Unmarshal([]byte(value.(string)), data)
	default:
		return fmt.Errorf("val type is valid, is %+v", value)
	}
}

// for valuer helper
func value(data interface{}) (interface{}, error) {
	vi := reflect.ValueOf(data)
	// 判断是否为 0 值
	if vi.IsZero() {
		return nil, nil
	}
	return json.Marshal(data)
}
