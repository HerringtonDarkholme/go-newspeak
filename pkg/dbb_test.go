package newspeak

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HerringtonDarkholme/go-newspeak/pkg/col"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserTable struct {
	Name col.StringColumn
	Age  col.IntColumn
}

func (u *UserTable) ToTableText() string {
	return "test"
}

func TestQueryGeneration(t *testing.T) {
	dbb, mock := newDBB(t)
	expectSQl := func(db *DBB, sql string) {
		// quotemeta is required since sqlmock is in regex mode
		mock.ExpectQuery(regexp.QuoteMeta(sql)).
			WillReturnRows(sqlmock.NewRows(nil))
		_ = dbb.Find(map[string]interface{}{})
	}
	u := &UserTable{
		Name: col.NewStr("name"),
		Age:  col.NewInt("age"),
	}

	expectSQl(dbb.Table(u), "SELECT * FROM `test`")
	expectSQl(dbb.Table(u).Select(u.Name), "SELECT name FROM `test`")
	expectSQl(dbb.Table(u).Select(u.Age), "SELECT age FROM `test`")

	// sub := dbb.db.Raw("avg(we)")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}

func newDBB(t *testing.T) (*DBB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		return nil, nil
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))
	if err != nil {
		t.Fatalf("an error '%s' happened during create gorm db", err)
		return nil, nil
	}
	dbb := New(gormDB)
	return dbb, mock
}
