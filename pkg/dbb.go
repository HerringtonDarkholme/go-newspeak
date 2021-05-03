package newspeak

import (
	"gorm.io/gorm"
)

type DBB struct {
	db *gorm.DB
}

func (dbb *DBB) Table(tables ...TableSpecifier) *DBB {
	return dbb
}

func (dbb *DBB) Select(expr ...Expression) *DBB {
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

func (dbb *DBB) Join(tables ...TableSpecifier) *DBB {
	return dbb
}

func (dbb *DBB) On(conditions ...WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Union(tables ...WhereCondition) *DBB {
	return dbb
}

func (dbb *DBB) Find(users interface{}) error {
	return nil
}
