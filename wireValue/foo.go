package wireValue

import "github.com/google/wire"

// 结构体定义
type Foo struct {
	X int
	Y int `wire:"-"`
}

// 供给者集合（顺序、绑定）
var WireSet = wire.Value(Foo{X: 2024})
