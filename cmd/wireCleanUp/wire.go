//go:build wireinject

package main

import (
	"github.com/google/wire"
	"wireDemo/wireCleanUp"
)

func NewWire() (string, func(), error) {
	wire.Build(wireCleanUp.WireSet)
	return "", func() {}, nil
}
