package db

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type QuerySeter interface {
	orm.QuerySeter
}

type DB interface {
	orm.Ormer
	From(table string) *Querier
	Execute(sql string, params ...interface{}) (int64, error)
}

type db struct {
	orm.Ormer
}

func NewDB() DB {
	o := orm.NewOrm()
	d := new(db)
	d.Ormer = o
	return d
}

func (d db) From(table string) *Querier {
	query := NewQuery(d, NewGenericSQLBuilder())
	query.From(table)
	return query
}

func (d db) Execute(sql string, params ...interface{}) (int64, error) {

	result, err := d.Raw(sql, params).Exec()
	sql = strings.TrimLeft(sql, " ")
	sql = strings.ToLower(sql)
	var isInsert = strings.HasPrefix(sql, "insert")
	if err != nil {
		return -1, err
	}
	if isInsert {
		return result.LastInsertId()
	} else {
		return result.RowsAffected()
	}
}

