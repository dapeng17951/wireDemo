package wireBaseDemo

import "github.com/google/wire"

var WireSet = wire.NewSet(NewFoo, NewBar, NewBaz)
