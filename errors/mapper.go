package errors

import (
	"errors"
	"maps"
)

type Mapper struct {
	Errors     map[error]any
	onNotFound any
}

func New(errors map[error]any, onNotFound any) *Mapper {
	return &Mapper{Errors: errors, onNotFound: onNotFound}
}

func NewEmpty() *Mapper {
	return New(make(map[error]any), nil)
}

func (m *Mapper) Get(err error) (out any, outErr error) {
	unwrapped := err
	for errors.Unwrap(unwrapped) != nil {
		unwrapped = errors.Unwrap(unwrapped)
	}

	out = m.onNotFound
	outErr = err

	// todo, find better way to deal with 'unhashable errors'
	defer func() { recover() }()

	foundErr, ok := m.Errors[unwrapped]
	if ok {
		return foundErr, err
	}

	return
}

type MapperBuilder struct {
	mapper *Mapper
}

func NewBuilder() *MapperBuilder {
	return &MapperBuilder{
		mapper: NewEmpty(),
	}
}

func (b *MapperBuilder) OnNotFound(onNotFound any) *MapperBuilder {
	b.mapper.onNotFound = onNotFound
	return b
}

func (b *MapperBuilder) AppendMap(errors ...map[error]any) *MapperBuilder {
	for _, err := range errors {
		maps.Copy(b.mapper.Errors, err)
	}
	return b
}

func (b *MapperBuilder) Build() *Mapper {
	return b.mapper
}
