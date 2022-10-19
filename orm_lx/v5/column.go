package v5

type Column struct {
	name string
}

func (c Column) expr() {}

type value struct {
	val any
}

func (v value) expr() {}

func valueOf(v any) value {
	return value{
		val: v,
	}
}

func C(name string) Column {
	return Column{
		name: name,
	}
}

// EQ 例如 C("id").Eq(12)
func (c Column) EQ(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    opEQ,
		right: exprOf(arg),
	}
}

func (c Column) LT(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    opLT,
		right: exprOf(arg),
	}
}

func (c Column) GT(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    opGT,
		right: exprOf(arg),
	}
}