package orm

import "database/sql"

type Database interface {
	Connect() error
	Query(clause string, args ...any) (results []map[string]any, headers []*sql.ColumnType, err error)
	Mutate(clause string, args ...any) (rowsAffected int64, err error)
}

type database struct {
	Driver     string
	DataSource string
	//
	DB *sql.DB
}

func NewDatabase(driver, dataSource string) Database {
	return &database{
		Driver:     driver,
		DataSource: dataSource,
	}
}

func (o *database) Connect() (err error) {
	o.DB, err = sql.Open(o.Driver, o.DataSource)
	return
}

func (o *database) getClause(raw string) string {
	c := &Clause{o.Driver, raw}
	return c.Get()
}
