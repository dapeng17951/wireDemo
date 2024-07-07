//go:build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"wireDemo/wireBaseDemo"
)

// 直接依赖 供给者的 注入方式
func NewWire(ctx context.Context) (wireBaseDemo.Baz, error) {
	wire.Build(wireBaseDemo.NewFoo, wireBaseDemo.NewBar, wireBaseDemo.NewBaz)
	return wireBaseDemo.Baz{}, nil
}

// 依赖 供给者集合的 注入方式
func NewWireSet(ctx context.Context) (wireBaseDemo.Baz, error) {
	wire.Build(wireBaseDemo.WireSet)
	return wireBaseDemo.Baz{}, nil
}

// 依赖 供给者集合的 注入方式 与 Value 结合
func NewWireSet2(ctx context.Context) (wireBaseDemo.Baz, error) {
	wire.Build(wireBaseDemo.WireSet2)
	return wireBaseDemo.Baz{}, nil
}
