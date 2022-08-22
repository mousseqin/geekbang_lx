package reflect

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"reflect"
)

func IterateField(val any) {
	res, err := iterateFields(val)
	if err != nil {
		spew.Dump(err)
		return
	}
	for k, v := range res {
		spew.Dump(k, v)
	}
}

func SetField(entity any, field string, newVal any) error {
	val := reflect.ValueOf(entity)
	typ := val.Type()

	// 只能是一级指针，类似 *User
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}
	typ = typ.Elem()
	val = val.Elem()

	// 利用 type 来判断 field 在不在
	if _, found := typ.FieldByName(field); !found {
		return errors.New("字段不存在")
	}

	// 这个地方判断不出来 field 在不在
	fd := val.FieldByName(field)
	if !fd.CanSet() {
		return errors.New("不可修改字段")
	}

	fd.Set(reflect.ValueOf(newVal))
	return nil
}

// iterateFields 返回所有的字段名字
// in 只能是结构体，或者结构体指针，可以是多重指针
func iterateFields(in any) (map[string]any, error) {
	typ := reflect.TypeOf(in)
	val := reflect.ValueOf(in)

	// 处理指针，要拿到指针指向的东西
	// 这里我们综合考虑了多重指针的效果
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	// 如果不是结构体，就返回 error
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("非法类型")
	}

	num := typ.NumField()
	res := make(map[string]any, num)
	for i := 0; i < num; i++ {
		fd := typ.Field(i)
		fdVal := val.Field(i)
		if fd.IsExported() {
			res[fd.Name] = fdVal.Interface()
		} else {
			// 为了演示效果，不公开字段我们用零值来填充
			res[fd.Name] = reflect.Zero(fd.Type).Interface()
		}
	}
	return res, nil
}
