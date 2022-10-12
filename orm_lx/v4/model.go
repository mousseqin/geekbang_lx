package v4

type model struct {
	// tableName 结构体对应的表名
	tableName string
	fieldMap  map[string]*field
}

// field 字段
type field struct {
	colName string
}
