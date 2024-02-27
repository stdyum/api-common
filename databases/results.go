package databases

import "database/sql"

func AssertRowAffected(result sql.Result) error {
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func AssertRowAffectedErr(result sql.Result, err error) error {
	if err != nil {
		return err
	}

	return AssertRowAffected(result)
}
