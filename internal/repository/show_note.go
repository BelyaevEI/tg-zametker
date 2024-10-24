package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/BelyaevEI/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) ShowNotes(userID int64) ([]string, error) {

	notes := make([]string, 0)

	builder := sq.Select(noteColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).Where(sq.Eq{userIDColumn: userID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "show_notes",
		QueryRaw: query,
	}

	err = r.db.DB().ScanAllContext(context.Background(), &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repo) getNote(userID int64, noteText string) (time.Time, error) {
	var timer sql.NullTime

	builder := sq.Select(noteTimeColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).Where(sq.Eq{userIDColumn: userID}).Where(sq.Eq{noteColumn: noteText})

	query, args, err := builder.ToSql()
	if err != nil {
		return time.Time{}, err
	}

	q := db.Query{
		Name:     "get_note",
		QueryRaw: query,
	}

	err = r.db.DB().ScanOneContext(context.Background(), &timer, q, args...)
	if err != nil {
		return time.Time{}, err
	}

	if !timer.Valid {
		return time.Time{}, errors.New("null time")
	}

	return timer.Time, nil
}
