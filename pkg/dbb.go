package newspeak

import (
	"gorm.io/gorm"
)

type DBB struct {
	db     *gorm.DB
	Select SelectFn
	Order  OrderFn
	Union  UnionFn
	Join   JoinFn
}

type SelectFn func(...SelectExpr) *DBB
type OrderFn func(...Expression) *DBB
type UnionFn func(...TableSpecifier) *DBB
type JoinFn func(...TableSpecifier) *DBB

func (fn SelectFn) Distinct(exprs ...SelectExpr) *DBB {
	return fn(exprs...)
}
func (fn SelectFn) All(exprs ...SelectExpr) *DBB {
	return fn(exprs...)
}
func (fn SelectFn) DistinctRow(exprs ...SelectExpr) *DBB {
	return fn(exprs...)
}

func (fn OrderFn) Asc(exprs ...Expression) *DBB {
	return fn(exprs...)
}
func (fn OrderFn) Desc(exprs ...Expression) *DBB {
	return fn(exprs...)
}

func (fn UnionFn) All(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn UnionFn) Distinct(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}

func (fn JoinFn) Inner(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Outer(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Left(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Right(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Cross(tables ...TableSpecifier) *DBB {
	return fn(tables...)
}

func (dbb *DBB) Table(table TableSpecifier) *DBB {
	dbb.db = dbb.db.Table(table.ToTableText())
	return dbb
}

func (dbb *DBB) Where(conditions ...WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Group(expr ...Expression) *DBB {
	return dbb
}

func (dbb *DBB) Having(conditions ...WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Window(specs ...WindowSpec) *DBB {
	return dbb
}

func (dbb *DBB) Limit(limit int64) *DBB {
	return dbb
}

func (dbb *DBB) Offset(offset int64) *DBB {
	return dbb
}

func (dbb *DBB) On(conditions ...WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Find(dest interface{}) error {
	return dbb.db.Find(dest).Error
}

// do not export this type, only string literal is permitted
type rawSql string

// an escape hatch for ultimatly complicated beast SQL
func (dbb *DBB) Query(sql rawSql, args ...interface{}) *DBB {
	return dbb
}
