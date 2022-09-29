package v1

import (
	"reflect"
	"strings"
)

type Selector[T any] struct {
	table string
}

func (s *Selector[T]) From(tbl string) *Selector[T] {
	s.table = tbl
	return s
}

func (s *Selector[T]) Build() (*Query, error) {
	var sb strings.Builder
	sb.WriteString("SELECT * FROM ")
	if s.table == "" {
		var t T
		sb.WriteString("`")
		sb.WriteString(reflect.TypeOf(t).Name())
		sb.WriteString("`")
	} else {
		sb.WriteString(s.table)
	}
	sb.WriteString(";")
	return &Query{
		SQL: sb.String(),
	}, nil
}

func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}
