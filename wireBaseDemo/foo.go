package wireBaseDemo

import "github.com/google/wire"

// Foo 结构体声明
type Foo struct {
	X int
}

// Foo 供给者
func NewFoo() Foo {
	return Foo{X: 2024}
}

// Foo 供给者，与 Value 结合
var WireFooSet = wire.Value(Foo{X: 2025})
