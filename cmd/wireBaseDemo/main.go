package main

import "context"

func main() {
	baz, err := NewWire(context.Background())
	if err != nil {
		panic(err)
	}
	println(baz.X)

	//与 value 结合的情况
	baz2, err := NewWireSet2(context.Background())

	if err != nil {
		panic(err)
	}
	println(baz2.X)
}
