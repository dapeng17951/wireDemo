package wireStruct

import "github.com/google/wire"

type Foo int
type Bar int

// Foo 供给者
func NewFoo() Foo {
	return 1
}

// Bar 供给者
func NewBar() Bar {
	return 2
}

// 结构体定义
type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

// 供给者 集合
var WireSet = wire.NewSet(NewFoo, NewBar, wire.Struct(new(FooBar), "MyFoo", "MyBar"))

// 供给者 单个
var WireSetTwo = wire.NewSet(NewFoo, NewBar, wire.Struct(new(FooBar), "MyFoo"))
