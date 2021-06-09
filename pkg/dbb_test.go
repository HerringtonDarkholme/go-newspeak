package newspeak

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestQueryGeneration(t *testing.T) {
	dbb, mock := newDBB(t)

	mock.MatchExpectationsInOrder(false)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `test`")).
		WillReturnRows(sqlmock.NewRows(nil))

	// sub := dbb.db.Raw("avg(we)")
	dbb.db.Debug().Table("test").Find(map[string]interface{}{})

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
	dbb := &DBB{
		db: gormDB,
	}
	return dbb, mock
}
