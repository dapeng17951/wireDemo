//go:build wireinject
// +build wireinject

package wireBaseDemo

import (
	"context"
	"github.com/google/wire"
)

// 注入
func NewWire(ctx context.Context) (Baz, error) {
	wire.Build(NewFoo, NewBar, NewBaz)
	return Baz{}, nil
}

//在同包wireBaseDemo下，既然提示undefined
//在main包下，既然可以正常generate
/*
wire: C:\Users\32763\Desktop\workspace\golang\wireDemo\wireBaseDemo\wire.go:12:36: undefined: Baz
wire: C:\Users\32763\Desktop\workspace\golang\wireDemo\wireBaseDemo\wire.go:13:13: undefined: NewFoo
wire: C:\Users\32763\Desktop\workspace\golang\wireDemo\wireBaseDemo\wire.go:13:21: undefined: NewBar
wire: C:\Users\32763\Desktop\workspace\golang\wireDemo\wireBaseDemo\wire.go:13:29: undefined: NewBaz
wire: C:\Users\32763\Desktop\workspace\golang\wireDemo\wireBaseDemo\wire.go:14:9: undefined: Baz
*/
