package wireBaseDemo

import "github.com/google/wire"

var WireSet = wire.NewSet(NewFoo, NewBar, NewBaz)

// 与 value 结合
var WireSet2 = wire.NewSet(WireFooSet, NewBar, NewBaz)
