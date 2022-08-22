package reflect

import (
	"errors"
	"reflect"
)

// Iterate 迭代数组，切片，或者字符串
func Iterate(input any) ([]any, error) {
	val := reflect.ValueOf(input)
	typ := val.Type()
	kind := typ.Kind()
	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		return nil, errors.New("非法类型")
	}

	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		res = append(res, ele.Interface())
	}
	return res, nil
}

// IterateMapV1 返回键，值
func IterateMapV1(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("非法类型")
	}
	kr := make([]any, 0, val.Len())
	vr := make([]any, 0, val.Len())
	for _, k := range val.MapKeys() {
		kr = append(kr, k.Interface())
		v := val.MapIndex(k)
		vr = append(vr, v.Interface())
	}

	return kr, vr, nil
}

// IterateMapV2 返回键，值
func IterateMapV2(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("非法类型")
	}
	kr := make([]any, 0, val.Len())
	vr := make([]any, 0, val.Len())
	itr := val.MapRange()
	for itr.Next() {
		kr = append(kr, itr.Key().Interface())
		vr = append(vr, itr.Value().Interface())
	}

	return kr, vr, nil
}
