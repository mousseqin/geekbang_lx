package kratos

import (
	"context"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type ctxKey1 struct{}
	type ctxKey2 struct{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ctx1 := context.WithValue(context.Background(), ctxKey1{}, "https://github.com/go-kratos/")
	ctx2 := context.WithValue(ctx, ctxKey2{}, "https://go-kratos.dev/")

	ctx, cancel = Merge(ctx1, ctx2)
	defer cancel()

	got := ctx.Value(ctxKey1{})
	value1, ok := got.(string)
	if !ok {
		t.Errorf("expect %v, got %v", true, ok)
	}
	if !reflect.DeepEqual(value1, "https://github.com/go-kratos/") {
		t.Errorf("expect %v, got %v", "https://github.com/go-kratos/", value1)
	}

	got2 := ctx.Value(ctxKey2{})
	value2, ok := got2.(string)
	if !ok {
		t.Errorf("expect %v, got %v", true, ok)
	}
	if !reflect.DeepEqual(value2, "https://go-kratos.dev/") {
		t.Errorf("expect %v, got %v", " https://go-kratos.dev/", value2)
	}

	t.Log(ctx)
}