package log

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/farismfirdaus/stockbit-technical-test/microservices/logging/model"
	repository "github.com/farismfirdaus/stockbit-technical-test/microservices/logging/repository"
)

type log struct {
	DB *sql.DB
}

func BuildLog(db *sql.DB) repository.LogInterface {
	return &log{
		DB: db,
	}
}

func (l *log) Insert(data *model.DbLog) error {
	query := sq.Insert("log").
		Columns("timestamp", "method", "request", "response").
		Values(sq.Expr("NOW()"), data.Method, data.Request, data.Response).
		Suffix("RETURNING id").
		RunWith(l.DB).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRow().Scan(&data.ID)
	if err != nil {
		return err
	}
	if data.ID == 0 {
		return errors.New("inserted id: 0")
	}
	return nil
}
