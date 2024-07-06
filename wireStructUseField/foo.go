package wireStructUseField

type Foo struct {
	S string
	N int
	F float64
}

// 属性 get
func GetS(foo Foo) string {
	// Bad! Use wire.FieldsOf instead.
	return foo.S
}

// 供给者
func NewFoo() Foo {
	return Foo{S: "Hello, World!", N: 1, F: 3.14}
}
