package v3

type Column struct {
	name string
}

func (c Column) expr() {}

type Value struct {
	val any
}

func (c Value) expr() {}

func valueOf(v any) Value {
	return Value{
		val: v,
	}
}

func C(n string) Column {
	return Column{
		name: n,
	}
}

// EQ 例如 C("id").Eq(12)
func (c Column) EQ() {

}
