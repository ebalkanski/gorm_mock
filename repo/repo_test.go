package repo_test

import (
	"database/sql"
	"gorm/repo"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSel(t *testing.T) {
	var conn *sql.DB
	var mock sqlmock.Sqlmock
	var err error

	conn, mock, err = sqlmock.New() // mock sql.DB
	assert.Nil(t, err)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})
	assert.Nil(t, err)

	r := repo.New(db)

	query := `SELECT`
	id := "abcdef"

	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(id)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)
	user, err := r.Select()
	assert.Nil(t, err)
	assert.Equal(t, &repo.User{ID: id}, user)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}

func TestIns(t *testing.T) {
	var conn *sql.DB
	var mock sqlmock.Sqlmock
	var err error

	conn, mock, err = sqlmock.New()
	assert.Nil(t, err)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	assert.Nil(t, err)

	r := repo.New(db)

	// select
	querySel := `SELECT`

	rowsSel := sqlmock.
		NewRows([]string{"id"}).
		AddRow("xxxxxx")

	mock.ExpectQuery(regexp.QuoteMeta(querySel)).
		WillReturnRows(rowsSel)

	// insert
	queryIns := `INSERT`
	id := "yyyyyy"

	mock.ExpectExec(regexp.QuoteMeta(queryIns)).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = r.Insert(id)
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}
