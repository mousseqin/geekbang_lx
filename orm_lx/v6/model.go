package v6

import "geekbang_lx/orm_lx/v6/internal/errs"

type ModelOpt func(model *Model) error

type Model struct {
	// tableName 结构体对应的表名
	tableName string
	fieldMap  map[string]*field
}

func ModelWithTableName(tableName string) ModelOpt {
	return func(model *Model) error {
		model.tableName = tableName
		return nil
	}
}

func ModelWithColumnName(field, columnName string) ModelOpt {
	return func(model *Model) error {
		fd, ok := model.fieldMap[field]
		if !ok {
			return errs.NewErrUnknownField(field)
		}
		fd.colName = columnName
		return nil
	}
}

// field 字段
type field struct {
	colName string
}

// 我们支持的全部标签上的 key 都放在这里
// 方便用户查找，和我们后期维护
const (
	tagKeyColumn = "column"
)

// 用户自定义一些模型信息的接口，集中放在这里
// 方便用户查找和我们后期维护

// TableName 用户实现这个接口来返回自定义的表名
type TableName interface {
	TableName() string
}
