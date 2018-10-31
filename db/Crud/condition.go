package db

import (
	"errors"
	"zerogo/utils"
)

type ConditionType int

const (
	SEGMENT ConditionType = 1 + iota
	EQ
	NOT_EQ
	IN
	NOT_IN
	LIKE
	NOT_LIKE
	BETWEEN
	NOT_BETWEEN
	LESS
	LE
	GREAT
	GE
	NULL
	NOT_NULL
)

var conditionTypes = [...]string{
	"",
	"=",
	"!=",
	"",
	"",
	"like",
	"not like",
	"",
	"",
	"<",
	"<=",
	">",
	">=",
	"",
	"",
}

func (c ConditionType) String() string {
	return conditionTypes[c - 1]
}

type Condition struct {
	Type   ConditionType
	Column string
	Params []interface{}
}

func NewCondition(condition ConditionType, column string, param interface{}) Condition {
	var c Condition
	c.Type = condition
	c.Column = column
	c.Params = utils.ToSlice(param)
	return c
}

func (c Condition) ToSQL(sqlBuilder SQLBuilder) (string, error) {
	if len(c.Type.String()) > 0 {
		return sqlBuilder.EscapeColumn(c.Column) + " " + c.Type.String() + " ?", nil
	}
	if len(c.Column) < 1 {
		return "", errors.New("c.Column is empty")
	}
	switch c.Type {
	case SEGMENT:
		return "(" + c.Column + ")", nil
	case IN, NOT_IN:
		var sql = sqlBuilder.EscapeColumn(c.Column)
		if c.Type == NOT_IN {
			sql += " not"
		}
		sql += " in ("
		for i, _ := range c.Params {
			if (i > 0) {
				sql += ", "
			}
			sql += "?"
		}
		sql += ")"
		return sql, nil
	case BETWEEN, NOT_BETWEEN:
		var sql = sqlBuilder.EscapeColumn(c.Column)
		if c.Type == NOT_BETWEEN {
			sql += " not "
		}
		sql += " between ? and ? "
		return sql, nil
	case NULL:
		return sqlBuilder.EscapeColumn(c.Column) + " is null", nil;
	case NOT_NULL:
		return sqlBuilder.EscapeColumn(c.Column) + " is not null", nil;
	}
	return "", errors.New("sql analysis error")
}
