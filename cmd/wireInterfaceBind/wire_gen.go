// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wireDemo/wireInterfaceBind"
)

// Injectors from wire.go:

func NewWire() string {
	myFooer := wireInterfaceBind.NewMyFooer()
	string2 := wireInterfaceBind.NewBar(myFooer)
	return string2
}
