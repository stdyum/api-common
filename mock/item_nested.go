package mock

import (
	"context"
)

type DataItemNested[D any, I any, P any] struct {
	Name          string
	Config        ConfigItem
	Data          D
	Generate      func(ctx context.Context, i int, data *D, previous P) (I, error)
	GenerateItems func(ctx context.Context, data *D, previous P) ([]I, error)
	Insert        func(ctx context.Context, i int, items []I, previous P) error
	Nested        []RawItem
}

func (item DataItemNested[D, I, P]) Build() RawItem {
	return RawItem{
		Name:          item.Name,
		Config:        item.Config,
		Data:          item.Data,
		Generate:      item.convertGenerateFunc(),
		GenerateItems: item.convertGenerateItemsFunc(),
		Insert:        item.convertInsertFunc(),
		Nested:        item.Nested,
	}
}

func (item DataItemNested[D, I, P]) convertGenerateFunc() func(ctx context.Context, i int, data *any, previous any) (any, error) {
	if item.Generate == nil {
		return nil
	}

	return func(ctx context.Context, i int, data *any, previous any) (any, error) {
		var d *D
		if data != nil && *data != nil {
			dd := (*data).(D)
			d = &dd
		}

		var p P
		if previous != nil {
			p = previous.(P)
		}

		return item.Generate(ctx, i, d, p)
	}
}

func (item DataItemNested[D, I, P]) convertGenerateItemsFunc() func(ctx context.Context, data *any, previous any) ([]any, error) {
	if item.GenerateItems == nil {
		return nil
	}

	return func(ctx context.Context, data *any, previous any) ([]any, error) {
		var d *D
		if data != nil && *data != nil {
			dd := (*data).(D)
			d = &dd
		}

		var p P
		if previous != nil {
			p = previous.(P)
		}

		items, err := item.GenerateItems(ctx, d, p)
		if err != nil {
			return nil, err
		}

		iItems := make([]any, len(items))
		for index, i := range items {
			iItems[index] = i
		}

		return iItems, nil
	}
}

func (item DataItemNested[D, I, P]) convertInsertFunc() func(ctx context.Context, i int, iItems []any, previous any) error {
	if item.Insert == nil {
		return nil
	}

	return func(ctx context.Context, i int, iItems []any, previous any) error {
		items := make([]I, len(iItems))
		for index, i := range iItems {
			items[index] = i.(I)
		}

		var p P
		if previous != nil {
			p = previous.(P)
		}

		return item.Insert(ctx, i, items, p)
	}
}
