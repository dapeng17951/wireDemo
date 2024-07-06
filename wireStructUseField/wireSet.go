package wireStructUseField

import "github.com/google/wire"

var WireSet = wire.NewSet(NewFoo, GetS)
var WireSetTwo = wire.NewSet(NewFoo, wire.FieldsOf(new(Foo), "S"))
