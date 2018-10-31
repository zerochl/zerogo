package db

import (
	_ "github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego"
	"fmt"
)

type Querier struct {
	DB
	SQLBuilder
	table string
}

func NewQuery(db DB, sqlBuilder SQLBuilder) *Querier {
	q := new(Querier)
	q.DB = db
	q.SQLBuilder = sqlBuilder
	return q
}

func (q *Querier) From(table string) *Querier {
	q.table = table
	q.SQLBuilder.SetTable(table)
	return q
}

func (q *Querier) Table() string {
	return q.table
}

func (q *Querier) GroupBy(groupBy string) *Querier {
	q.SQLBuilder.GroupBy(groupBy)
	return q
}
func (q *Querier) Limit(offset int, rowCount int) *Querier {
	q.SQLBuilder.Limit(offset, rowCount)
	return q
}

func (q *Querier) OrderBy(orderBy string) *Querier {
	if len(orderBy) == 0 {
		return q
	}
	q.SQLBuilder.OrderBy(orderBy)
	return q
}

func (q *Querier) Select(columns ... string) *Querier {
	q.SQLBuilder.Select(columns)
	return q
}

func (q *Querier) Join(join string) *Querier {
	q.SQLBuilder.Join(join)
	return q
}

func (q *Querier) where(c Condition) *Querier {
	sql, err := c.ToSQL(q.SQLBuilder)
	if (nil == err) {
		q.SQLBuilder.Where(sql, c.Params)
	} else {
		beego.Error(err)
	}
	return q
}

func (q *Querier) Segment(sql string, params ...interface{}) *Querier {
	return q.where(NewCondition(SEGMENT, sql, params))
}

func (q *Querier) Eq(name string, value interface{}) *Querier {
	return q.where(NewCondition(EQ, name, value))
}

func (q *Querier) Where(name string, value interface{}) *Querier {
	return q.where(NewCondition(EQ, name, value))
}

func (q *Querier) Not(name string, value interface{}) *Querier {
	return q.where(NewCondition(NOT_EQ, name, value))
}

func (q *Querier) In(name string, values []interface{}) *Querier {
	return q.where(NewCondition(IN, name, values))
}

func (q *Querier) NotIn(name string, values []interface{}) *Querier {
	return q.where(NewCondition(NOT_IN, name, values))
}
func (q Querier) Between(name string, values []interface{}) *Querier {
	return q.where(NewCondition(BETWEEN, name, values))
}

func (q *Querier) NotBetween(name string, values []interface{}) *Querier {
	return q.where(NewCondition(NOT_BETWEEN, name, values))
}

func (q *Querier) Less(name string, value interface{}) *Querier {
	return q.where(NewCondition(LESS, name, value))
}

func (q *Querier) LessOrEquals(name string, value interface{}) *Querier {
	return q.where(NewCondition(LE, name, value))
}

func (q *Querier) Great(name string, value interface{}) *Querier {
	return q.where(NewCondition(GREAT, name, value))
}

func (q *Querier) GreatOrEquals(name string, value interface{}) *Querier {
	return q.where(NewCondition(GE, name, value))
}

func (q *Querier) IsNull(name string, value interface{}) *Querier {
	return q.where(NewCondition(NULL, name, value))
}

func (q *Querier) IsNotNull(name string, value interface{}) *Querier {
	return q.where(NewCondition(NOT_NULL, name, value))
}

func (q Querier) Like(name string, value interface{}) *Querier {
	return q.where(NewCondition(LIKE, name, value))
}

func (q *Querier) NotLike(name string, value interface{}) *Querier {
	return q.where(NewCondition(NOT_LIKE, name, value))
}

func (q *Querier) ToSql() string {
	return q.SQLBuilder.ToSql()
}

func (q *Querier) First(container interface{}) error {
	var err error
	var sql = q.ToSql()
	err = q.Raw(sql, q.Parameters()).QueryRow(container)

	return err
}

func (q *Querier) All(container interface{}) error {
	var err error
	var sql = q.ToSql()
	_, err = q.Raw(sql, q.Parameters()).QueryRows(container)
	return err
}

func (q *Querier) Pagination(container interface{}, page int, pageSize int) (*Pagination, error) {
	var err error
	var totalItem int
	var hasNext bool
	q.Limit((page - 1) * pageSize, pageSize)
	q.Raw(q.ToCountSql(), q.Parameters()).QueryRow(&totalItem)
	var sql = q.ToSql()

	_, err = q.Raw(sql, q.Parameters()).QueryRows(container)

	pagination := NewPagination(page, totalItem, hasNext)
	pagination.setPerPage(pageSize)
	pagination.hasNext = pagination.TotalPages() > page
	pagination.SetData(container)

	return pagination, err
}

func (q *Querier) FillPagination(container interface{}, pagination *Pagination) (int64, error) {

	var err error
	var count int64
	var totalItem int
	var page = pagination.Page
	var pageSize = pagination.PerPage

	q.Limit((page - 1) * pageSize, pageSize)
	var sql = q.ToCountSql()
	fmt.Print("++++++++++++++++++++", sql)
	q.Raw(sql, q.Parameters()).QueryRow(&totalItem)

	sql = q.ToSql()
	count, err = q.Raw(sql, q.Parameters()).QueryRows(container)

	pagination.Total = totalItem
	pagination.hasNext = pagination.TotalPages() > page
	pagination.SetData(container)

	return count, err
}
