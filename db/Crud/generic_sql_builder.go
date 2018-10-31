package db

import (
	"strings"
	"strconv"
	"github.com/astaxie/beego"
)

type GenericSQLBuilder struct {
	Table      string
	Columns    []string
	joins      []string
	conditions []string
	GroupBys   []string
	OrderBys   []string
	Offset     int
	RowCount   int
	parameters []interface{}
}

func NewGenericSQLBuilder() *GenericSQLBuilder {
	builder := new(GenericSQLBuilder)
	builder.Offset = -1
	return builder
}

func (c *GenericSQLBuilder) SetTable(table string) {
	c.Table = table
}

func (c *GenericSQLBuilder) ToCountSql() string {
	if len(c.GroupBys) > 0 {
		var sql = string("select ")
		sql += " count(*) from " + c.Table + c.toJoinSql()
		if len(c.conditions) > 0 {
			sql += " where "
			for i, condition := range c.conditions {
				if (i > 0) {
					sql += " and "
				}
				sql += condition
			}
		}
		return sql

	}
	return "select count(*) from (" + c.toSQL0(true, false, false) + ") a";
}
func (c *GenericSQLBuilder) ToSql() string {
	return c.toSQL0(true, true, true);
}
func (c *GenericSQLBuilder) toSQL0(withGroup bool, withOrderBy bool, withLimit bool) string {
	var sql = string("select ")
	sql += c.toSelectColumns() + " from " + c.Table
	if len(c.conditions) > 0 {
		sql += " where "
		for i, condition := range c.conditions {
			if (i > 0) {
				sql += " and "
			}
			sql += condition
		}
	}

	beego.Debug("conditions", c.conditions, "parameters", c.parameters, "offset", c.Offset, "rowCount", c.RowCount)
	if (withGroup) {
		sql += c.toGroupBySql()
	}
	if (withOrderBy) {
		sql += c.toOrderBySQL()
	}
	if (withLimit) {
		sql += c.toLimitSQL()
	}
	return sql
}
func (c *GenericSQLBuilder) toSelectColumns() string {
	if len(c.Columns) == 0 {
		return "*"
	}
	return strings.Join(c.Columns, ",")
}

func (c *GenericSQLBuilder) toJoinSql() string {
	if len(c.joins) == 0 {
		return ""
	}
	return "\n" + strings.Join(c.joins, "\n")
}

func (c *GenericSQLBuilder) toGroupBySql() string {
	if len(c.GroupBys) == 0 {
		return ""
	}
	return " group by " + strings.Join(c.GroupBys, ",")
}

func (c *GenericSQLBuilder) toOrderBySQL() string {
	if len(c.OrderBys) == 0 {
		return ""
	}
	return " order by " + strings.Join(c.OrderBys, ",")
}

func (c *GenericSQLBuilder) toLimitSQL() string {
	if c.Offset < 0 {
		return ""
	}
	return " limit " + strconv.Itoa(c.Offset) + ", " + strconv.Itoa(c.RowCount)
}

func (c *GenericSQLBuilder) Where(condition string, params ... interface{}) {
	c.conditions = append(c.conditions, condition)
	c.parameters = append(c.parameters, params)
}

func (c *GenericSQLBuilder) ClearSelect() {
	c.Columns = nil
}

func (c *GenericSQLBuilder) Select(columns [] string) {
	for _, column := range columns {
		c.Columns = append(c.Columns, column)
	}
}

func (c *GenericSQLBuilder) OrderBy(orderBy string) {
	c.OrderBys = append(c.OrderBys, orderBy)
}

func (c *GenericSQLBuilder) Join(join string) {
	c.joins = append(c.joins, join)
}

func (c *GenericSQLBuilder) GroupBy(groupBy string) {
	c.GroupBys = append(c.GroupBys, groupBy)
}

func (c *GenericSQLBuilder) Limit(offset int, rowCount int) {
	c.Offset = offset
	c.RowCount = rowCount
}

func (c *GenericSQLBuilder) HasLimit() bool {
	return c.Offset > 0
}

func (c *GenericSQLBuilder) EscapeColumn(column string) string {
	return column
}

func (c *GenericSQLBuilder) Parameters() []interface{} {
	return c.parameters
}

func (c *GenericSQLBuilder) Conditions() []string {
	return c.conditions
}



