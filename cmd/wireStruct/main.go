package main

func main() {
	bar := NewWire()
	println(bar.MyBar)
	println(bar.MyFoo)

	bar2 := NewWireTwo()

	println(bar2.MyFoo)

}
