package wireInterfaceBind

type Bar string

// 供给者，依赖注入接口
func NewBar(foo Fooer) string {
	return foo.Foo()
}
