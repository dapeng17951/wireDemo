//go:build wireinject

package main

import (
	"github.com/google/wire"
	"wireDemo/wireStructUseField"
)

// 注方式1：get方法获取属性
func NewWire() string {
	wire.Build(wireStructUseField.WireSet)
	return ""
}

// 注方式2：直接属性
func NewWireTwo() string {
	wire.Build(wireStructUseField.WireSetTwo)
	return ""
}
