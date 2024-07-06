//go:build wireinject

package main

import (
	"github.com/google/wire"
	"wireDemo/wireStruct"
)

// 注入
func NewWire() wireStruct.FooBar {
	wire.Build(wireStruct.WireSet)
	return wireStruct.FooBar{}
}

func NewWireTwo() wireStruct.FooBar {
	wire.Build(wireStruct.WireSetTwo)
	return wireStruct.FooBar{}
}
