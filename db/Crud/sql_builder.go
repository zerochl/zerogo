package db

type SQLBuilder interface {
	SetTable(tableName string)
	Parameters() []interface{}
	ToCountSql() string
	ToSql() string
	Where(condition string, params ... interface{})
	Select(columns []string)
	ClearSelect()
	Join(joinSql string)
	GroupBy(groupBy string)
	OrderBy(orderBy string)
	Limit(offset int, rowCount int)
	HasLimit() bool
	EscapeColumn(column string) string
}