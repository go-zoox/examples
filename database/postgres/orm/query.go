package orm

import (
	"database/sql"
	"fmt"
	"time"
)

func (o *database) Query(clause string, args ...any) (results []map[string]any, headers []*sql.ColumnType, err error) {
	clauseX := o.getClause(clause)

	stmt, err := o.DB.Prepare(clauseX)
	if err != nil {
		return nil, nil, fmt.Errorf("prepare error: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query: %v", err)
	}

	headers, _ = rows.ColumnTypes()
	values := make([]interface{}, len(headers))
	for i, header := range headers {
		valueType := header.ScanType().String()
		switch valueType {
		case "string":
			values[i] = string("")
		case "bool":
			values[i] = bool(false)
		case "int64":
			values[i] = int64(0)
		case "time.Time":
			values[i] = time.Time{}
		default:
			values[i] = sql.RawBytes{}
		}
	}

	scans := make([]interface{}, len(headers))
	for i := range values {
		scans[i] = &values[i]
	}

	results = []map[string]interface{}{}
	i := 0
	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to iterate row: %v", err)
		}

		row := make(map[string]interface{})
		for j, v := range values {
			key := headers[j].Name()
			value := v
			row[key] = value
		}

		results = append(results, row)

		i++
	}

	err = rows.Err()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to scan rows: %v", err)
	}

	return results, headers, nil
}
