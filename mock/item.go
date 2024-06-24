package mock

import (
	"context"
)

type DataItem[D any, I any] struct {
	Name          string
	Config        ConfigItem
	Data          D
	Generate      func(ctx context.Context, i int, data *D) (I, error)
	GenerateItems func(ctx context.Context, data *D) ([]I, error)
	Insert        func(ctx context.Context, i int, items []I) error
	Nested        []RawItem
}

func (item DataItem[D, I]) Build() RawItem {
	return DataItemNested[D, I, any]{
		Name:          item.Name,
		Config:        item.Config,
		Data:          item.Data,
		Generate:      item.convertGenerateFunc(),
		GenerateItems: item.convertGenerateItemsFunc(),
		Insert:        item.convertInsertFunc(),
		Nested:        item.Nested,
	}.Build()
}

func (item DataItem[D, I]) convertGenerateFunc() func(ctx context.Context, i int, data *D, previous any) (I, error) {
	if item.Generate == nil {
		return nil
	}

	return func(ctx context.Context, i int, data *D, previous any) (I, error) {
		return item.Generate(ctx, i, data)
	}
}

func (item DataItem[D, I]) convertGenerateItemsFunc() func(ctx context.Context, data *D, previous any) ([]I, error) {
	if item.GenerateItems == nil {
		return nil
	}

	return func(ctx context.Context, data *D, previous any) ([]I, error) {
		return item.GenerateItems(ctx, data)
	}
}

func (item DataItem[D, I]) convertInsertFunc() func(ctx context.Context, i int, items []I, previous any) error {
	if item.Insert == nil {
		return nil
	}

	return func(ctx context.Context, i int, items []I, previous any) error {
		return item.Insert(ctx, i, items)
	}
}
