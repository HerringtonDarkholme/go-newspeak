package newspeak

import (
	"gorm.io/gorm"
)

type DBB struct {
	db *gorm.DB
}

func (dbb *DBB) Table() *DBB {
	return dbb
}

func (dbb *DBB) Select() *DBB {
	return dbb
}

func (dbb *DBB) Where() *DBB {
	return dbb
}

func (dbb *DBB) Group() *DBB {
	return dbb
}

func (dbb *DBB) Having() *DBB {
	return dbb
}

func (dbb *DBB) Window() *DBB {
	return dbb
}

func (dbb *DBB) Limit() *DBB {
	return dbb
}

func (dbb *DBB) Offset() *DBB {
	return dbb
}

func (dbb *DBB) Join() *DBB {
	return dbb
}

func (dbb *DBB) On() *DBB {
	return dbb
}

func (dbb *DBB) Union() *DBB {
	return dbb
}
