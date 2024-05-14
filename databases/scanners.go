package databases

import (
	"database/sql"

	"github.com/stdyum/api-common/errors"
)

type Scan interface {
	Scan(...any) error
}

type Scanner interface {
	Scan
	Next() bool
}

func ScanArray[T any](scanner Scanner, rowScan func(row Scan) (T, error)) ([]T, error) {
	var rows []T
	for scanner.Next() {
		row, err := rowScan(scanner)
		if err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	if len(rows) == 0 {
		return nil, sql.ErrNoRows
	}

	return rows, nil
}

func ScanArrayErr[T any](scanner Scanner, rowScan func(row Scan) (T, error), err error) ([]T, error) {
	if err != nil {
		return nil, err
	}

	return ScanArray(scanner, rowScan)
}

func ScanPagination[T any](scanner Scanner, rowScan func(row Scan) (T, error), amount int) ([]T, int, error) {
	array, err := ScanArray(scanner, rowScan)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, amount, nil
	}

	return array, amount, err
}

func ScanPaginationErr[T any](scanner Scanner, rowScan func(row Scan) (T, error), amount int, err error) ([]T, int, error) {
	if err != nil {
		return nil, 0, err
	}

	return ScanPagination(scanner, rowScan, amount)
}
