package homework

import (
	"errors"
	"reflect"
	"strings"
)

type Options struct {
	Args   []interface{}
	strVal string
	strBld strings.Builder

	stuRepeat map[string]bool
}

var errInvalidEntity = errors.New("invalid entity")

// InsertStmt 作业里面我们这个只是生成 SQL，所以在处理 sql.NullString 之类的接口
// 只需要判断有没有实现 driver.Valuer 就可以了
func (o *Options) InsertStmt(entity interface{}) (string, []interface{}, error) {
	if entity == nil {
		return o.strVal, o.Args, errInvalidEntity
	}
	val := reflect.ValueOf(entity)
	typ := val.Type()

	// 处理指针，要拿到指针指向的东西
	// 这里我们综合考虑了多重指针的效果
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
		kind := typ.Kind()
		if kind == reflect.Pointer {
			return o.strVal, o.Args, errInvalidEntity
		}
	}
	// 只能是一级指针，类似 *User
	//if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
	if typ.Kind() != reflect.Struct {
		return o.strVal, o.Args, errInvalidEntity
	}
	num := typ.NumField()
	if num == 0 {
		return o.strVal, o.Args, errInvalidEntity
	}
	// 使用 strings.Builder 来拼接 字符串
	//o.strBld = strings.Builder{}
	// 构造 INSERT INTO XXX，XXX 是你的表名，这里我们直接用结构体名字
	o.strBld.WriteString("INSERT INTO " + "`" + typ.Name() + "`(")

	// 遍历所有的字段，构造出来的是 INSERT INTO XXX(col1, col2, col3)
	// 在这个遍历的过程中，你就可以把参数构造出来
	// 如果你打算支持组合，那么这里你要深入解析每一个组合的结构体
	// 并且层层深入进去
	o.Recursive(num, typ, val)

	// 拼接 VALUES，达成 INSERT INTO XXX(col1, col2, col3) VALUES
	// 再一次遍历所有的字段，要拼接成 INSERT INTO XXX(col1, col2, col3) VALUES(?,?,?)
	sbd := o.strBld.String()
	sbd = strings.TrimRight(sbd, ",")
	o.strBld.Reset()
	o.strBld.WriteString(sbd)
	o.strVal = strings.TrimRight(o.strVal, ",")
	// 注意，在第一次遍历的时候我们就已经拿到了参数的值，所以这里就是简单拼接 ?,?,?
	o.strBld.WriteString(") VALUES(" + o.strVal + ");")
	return o.strBld.String(), o.Args, nil
}

func (o *Options) Recursive(num int, typ reflect.Type, val reflect.Value) {
	for i := 0; i < num; i++ {
		if i == 0 {
			if bl := o.stuRepeat[typ.Name()]; bl {
				break
			}
		}
		fd := typ.Field(i)
		fdVal := val.Field(i)
		o.stuRepeat[typ.Name()] = true
		if fdVal.Kind() == reflect.Struct && fd.Anonymous {
			n := fdVal.NumField()
			o.Recursive(n, fd.Type, fdVal)
		} else {
			o.strBld.WriteString("`" + fd.Name + "`" + ",")
			o.Args = append(o.Args, fdVal.Interface())
			o.strVal = o.strVal + "?,"
		}
	}

}
