package repository

import (
	"context"
	"strconv"

	"github.com/BelyaevEI/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) DeleteNote(userID int64, noteNum string) error {
	notes, err := r.ShowNotes(userID)
	if err != nil {
		return err
	}
	indx, err := strconv.Atoi(noteNum)
	if err != nil {
		return err
	}

	note := notes[indx-1]

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{userIDColumn: userID}).Where(sq.Eq{noteColumn: note})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "note_delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(context.Background(), q, args...)
	if err != nil {
		return err
	}

	return nil
}
