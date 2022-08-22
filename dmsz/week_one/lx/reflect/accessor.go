package reflect

import (
	"errors"
	"reflect"
)

type Accessor struct {
	val reflect.Value
	typ reflect.Type
}

func NewReflectAccessor(val any) (*Accessor, error) {
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("invalid entity")
	}

	return &Accessor{
		val: reflect.ValueOf(val).Elem(),
		typ: typ.Elem(),
	}, nil
}

func (a *Accessor) Field(field string) (int, error) {
	if _, ok := a.typ.FieldByName(field); !ok {
		return 0, errors.New("非法字段")
	}
	return a.val.FieldByName(field).Interface().(int), nil
}

func (a *Accessor) SetField(field string, val int) error {
	if _, ok := a.typ.FieldByName(field); !ok {
		return errors.New("非法字段")
	}
	fdVal := a.val.FieldByName(field)
	if !fdVal.CanSet() {
		return errors.New("无法设置新值的字段")
	}
	fdVal.Set(reflect.ValueOf(val))
	return nil
}
