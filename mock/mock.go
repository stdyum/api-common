package mock

import (
	"context"
)

type Config struct {
	MaxSingleInsert int
}

type Mock struct {
	Config Config
	Item   RawItem
}

func (m *Mock) Mock() error {
	return m.MockContext(context.Background())
}

func (m *Mock) MockContext(ctx context.Context) error {
	return m.Item.Mock(ctx, m.Config)
}
