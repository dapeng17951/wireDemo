//go:build wireinject

package main

import (
	"github.com/google/wire"
	"wireDemo/wireInterfaceBind"
)

// 注入
func NewWire() string {
	wire.Build(wireInterfaceBind.WireSet)
	return ""
}
