package main

import "context"

func main() {
	baz, err := NewWire(context.Background())
	if err != nil {
		panic(err)
	}
	println(baz.X)
}
