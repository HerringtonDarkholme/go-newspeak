package newspeak

import (
	cond "github.com/HerringtonDarkholme/go-newspeak/pkg/conditions"
	"gorm.io/gorm"
)

type DBB struct {
	db     *gorm.DB
	Select SelectFn
	Order  OrderFn
	Union  UnionFn
	Join   JoinFn
}

func New(db *gorm.DB) *DBB {
	dbb := &DBB{
		db: db,
	}
	dbb.Select = func(exprs ...cond.SelectExpr) *DBB {
		selected := make([]string, 0, len(exprs))
		for _, expr := range exprs {
			selected = append(selected, expr.ToSelectText())
		}
		dbb.db.Select(selected)
		return dbb
	}
	return dbb
}

type SelectFn func(...cond.SelectExpr) *DBB
type OrderFn func(...cond.Expression) *DBB
type UnionFn func(...cond.TableSpecifier) *DBB
type JoinFn func(...cond.TableSpecifier) *DBB

func (fn SelectFn) Distinct(exprs ...cond.SelectExpr) *DBB {
	return fn(exprs...)
}
func (fn SelectFn) All(exprs ...cond.SelectExpr) *DBB {
	return fn(exprs...)
}
func (fn SelectFn) DistinctRow(exprs ...cond.SelectExpr) *DBB {
	return fn(exprs...)
}

func (fn OrderFn) Asc(exprs ...cond.Expression) *DBB {
	return fn(exprs...)
}
func (fn OrderFn) Desc(exprs ...cond.Expression) *DBB {
	return fn(exprs...)
}

func (fn UnionFn) All(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn UnionFn) Distinct(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}

func (fn JoinFn) Inner(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Outer(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Left(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Right(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}
func (fn JoinFn) Cross(tables ...cond.TableSpecifier) *DBB {
	return fn(tables...)
}

func (dbb *DBB) Table(table cond.TableSpecifier) *DBB {
	dbb.db = dbb.db.Table(table.ToTableText())
	return dbb
}

func (dbb *DBB) Where(conditions ...cond.WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Group(expr ...cond.Expression) *DBB {
	return dbb
}

func (dbb *DBB) Having(conditions ...cond.WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Window(specs ...cond.WindowSpec) *DBB {
	return dbb
}

func (dbb *DBB) Limit(limit int64) *DBB {
	return dbb
}

func (dbb *DBB) Offset(offset int64) *DBB {
	return dbb
}

func (dbb *DBB) On(conditions ...cond.WhereCondition) *DBB {
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
