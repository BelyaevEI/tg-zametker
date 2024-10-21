package repository

import (
	"context"

	"github.com/BelyaevEI/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) CreateNote(userID int64, note string) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIDColumn, noteColumn).
		Values(userID, note)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "create_note",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(context.Background(), q, args...)
	if err != nil {
		return err
	}

	return nil
}
