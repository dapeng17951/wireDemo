package main

func main() {
	s, f, err := NewWire()
	if err != nil {
		panic("failed")
	}
	println(s)
	defer f()
}
