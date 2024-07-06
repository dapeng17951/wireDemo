//go:build wireinject

package main

import (
	"github.com/google/wire"
	"wireDemo/wireValue"
)

// 注入方式1：通过set
func NewWire() wireValue.Foo {
	wire.Build(wireValue.WireSet)
	return wireValue.Foo{}
}

// 注入方式2：直接注入value
func NewWireTwo() wireValue.Foo {
	wire.Build(wire.Value(wireValue.Foo{X: 2025}))
	return wireValue.Foo{}
}
