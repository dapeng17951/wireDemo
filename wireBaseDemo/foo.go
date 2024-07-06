package wireBaseDemo

// Foo 结构体声明
type Foo struct {
	X int
}

// Foo 供给者
func NewFoo() Foo {
	return Foo{X: 2024}
}
