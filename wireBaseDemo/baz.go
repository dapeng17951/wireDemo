package wireBaseDemo

import (
	"context"
	"errors"
)

// Baz 结构提声明
type Baz struct {
	X int
}

// Baz 供给者，依赖注入 Bar,并在特定情况下，返回特定错误
// 这里的 context 会自动注入
func NewBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar in zero")
	}
	return Baz{
		X: bar.X,
	}, nil
}
