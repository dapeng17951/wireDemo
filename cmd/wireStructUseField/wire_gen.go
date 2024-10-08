// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wireDemo/wireStructUseField"
)

// Injectors from wire.go:

// 注方式1：get方法获取属性
func NewWire() string {
	foo := wireStructUseField.NewFoo()
	string2 := wireStructUseField.GetS(foo)
	return string2
}

// 注方式2：直接属性
func NewWireTwo() string {
	foo := wireStructUseField.NewFoo()
	string2 := foo.S
	return string2
}
