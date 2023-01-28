package orm

import (
	"fmt"
)

func (o *database) Mutate(clause string, args ...any) (rowsAffected int64, err error) {
	clauseX := o.getClause(clause)

	stmt, err := o.DB.Prepare(clauseX)
	if err != nil {
		return 0, fmt.Errorf("prepare error: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, fmt.Errorf("failed to query: %v", err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %v", err)
	}

	return rowCount, nil
}
