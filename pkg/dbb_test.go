package newspeak

import (
	"database/sql/driver"
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

type Args = []driver.Value

type testCase struct {
	dbb  *DBB
	sql  string
	args Args
}

func expectSQl(mock sqlmock.Sqlmock, tests []testCase) func(*testing.T) {
	for _, test := range tests {
		dbb, sql := test.dbb, test.sql
		// quotemeta is required since sqlmock is in regex mode
		mock.ExpectQuery(regexp.QuoteMeta(sql)).
			WithArgs(test.args...).
			WillReturnRows(sqlmock.NewRows(nil))
		_ = dbb.Find(map[string]interface{}{})
	}
	return func(t *testing.T) {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	}
}

func TestQueryGeneration(t *testing.T) {
	dbb, mock := newDBB(t)
	u := &UserTable{
		Name: col.NewStr("name"),
		Age:  col.NewInt("age"),
	}
	cases := map[string][]testCase{
		"Simple Query": {
			{dbb.Table(u), "SELECT * FROM `test`", nil},
		},
		"Select Field": {
			{dbb.Table(u).Select(u.Name), "SELECT name FROM `test`", nil},
			{dbb.Table(u).Select(u.Age), "SELECT age FROM `test`", nil},
		},
		"Simple Where": {
			{dbb.Table(u).Where(u.Name.Eq("Orwell")), "SELECT * FROM `test` WHERE name=?", Args{"Orwell"}},
		},
	}

	for testName, testCase := range cases {
		t.Run(testName, expectSQl(mock, testCase))
	}

	// sub := dbb.db.Raw("avg(we)")
}
