package wireInterfaceBind

type MyFooer string

// 接口实现
func (myFooer *MyFooer) Foo() string {
	return string(*myFooer)
}

// 供给者，反馈指针类型（接口实现时，指针类型）
func NewMyFooer() *MyFooer {
	fooer := new(MyFooer)
	*fooer = "hello wire"
	return fooer
}
