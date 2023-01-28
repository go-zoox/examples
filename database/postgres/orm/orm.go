package orm

import "database/sql"

type ORM interface {
	Database
}

type orm struct {
	db Database
}

func New(driver, dataSource string) ORM {
	return &orm{
		db: NewDatabase(driver, dataSource),
	}
}

func (o *orm) Connect() error {
	return o.db.Connect()
}

func (o *orm) Query(clause string, args ...any) (results []map[string]any, headers []*sql.ColumnType, err error) {
	return o.db.Query(clause, args...)
}

func (o *orm) Mutate(clause string, args ...any) (rowsAffected int64, err error) {
	return o.db.Mutate(clause, args...)
}
