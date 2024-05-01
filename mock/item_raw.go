package mock

import (
	"context"
)

type ConfigItem struct {
	Config
	Amount int
}

type RawItem struct {
	Config        ConfigItem
	Data          any
	Generate      func(ctx context.Context, i int, data *any, previous any) (any, error)
	GenerateItems func(ctx context.Context, data *any, previous any) ([]any, error)
	Insert        func(ctx context.Context, i int, items []any, previous any) error
	Nested        []RawItem
}

func (i *RawItem) Mock(ctx context.Context, config Config) error {
	return i.MockWithPrevious(ctx, config, nil)
}

func (i *RawItem) MockWithPrevious(ctx context.Context, config Config, previous any) error {
	items, err := i.getItems(ctx, config, previous)
	if err != nil {
		return err
	}

	if i.Insert != nil {
		if err = i.insert(ctx, config, items, previous); err != nil {
			return err
		}
	}

	for _, item := range items {
		for _, n := range i.Nested {
			if err = n.MockWithPrevious(ctx, config, item); err != nil {
				return err
			}
		}
	}

	return nil
}

func (i *RawItem) getItems(ctx context.Context, _ Config, previous any) ([]any, error) {
	if i.GenerateItems != nil {
		return i.GenerateItems(ctx, &i.Data, previous)
	}

	items := make([]any, i.Config.Amount)
	for index := range i.Config.Amount {
		item, err := i.Generate(ctx, index, &i.Data, previous)
		if err != nil {
			return nil, err
		}

		items[index] = item
	}

	return items, nil
}

func (i *RawItem) insert(ctx context.Context, config Config, items []any, previous any) error {
	var maxSingleInsert int
	if i.Config.MaxSingleInsert != 0 {
		maxSingleInsert = i.Config.MaxSingleInsert
	} else if config.MaxSingleInsert != 0 {
		maxSingleInsert = config.MaxSingleInsert
	}

	for index := 0; index < len(items); index += maxSingleInsert {
		end := index + maxSingleInsert
		if end > len(items) {
			end = len(items)
		}

		if err := i.Insert(ctx, index, items[index:end], previous); err != nil {
			return err
		}
	}

	return nil
}
