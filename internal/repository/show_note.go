package repository

import (
	"context"

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
