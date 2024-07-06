package wireBaseDemo

// Bar 结构体声明
type Bar struct {
	X int
}

// Bar 供给者，依赖注入了 Fool
func NewBar(foo Foo) Bar {
	return Bar{X: foo.X}
}
