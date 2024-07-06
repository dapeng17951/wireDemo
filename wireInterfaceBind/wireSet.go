package wireInterfaceBind

import "github.com/google/wire"

// 接口注入方式
// 需要指定接口绑定
var WireSet = wire.NewSet(NewBar, wire.Bind(new(Fooer), new(*MyFooer)), NewMyFooer)
