package repository

import (
	"context"

	"github.com/BelyaevEI/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) EditNote(userID, numberNote int64, note string) error {

	notes, err := r.ShowNotes(userID)
	if err != nil {
		return err
	}

	// т.к. выводим мы с 1 все заметки а индекс в слайсах с 0
	noteOld := notes[numberNote-1]

	//получим старое уведомление если есть
	timer, _ := r.getNote(userID, noteOld)

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{userIDColumn: userID}).Where(sq.Eq{noteColumn: noteOld})

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

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIDColumn, noteColumn, noteTimeColumn).
		Values(userID, note, timer)

	query, args, err = builder.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		Name:     "create_note",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(context.Background(), q, args...)
	if err != nil {
		return err
	}

	return nil
}
