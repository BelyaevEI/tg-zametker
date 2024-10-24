package repository

import "github.com/BelyaevEI/platform_common/pkg/db"

const (
	tableName      = "note"
	userIDColumn   = "user_id"
	noteColumn     = "note"
	noteTimeColumn = "note_time"
)

// Имплементация репо слоя
type Repositorer interface {
	CreateNote(userID int64, note string) error
	ShowNotes(userID int64) ([]string, error)
	DeleteNote(userID int64, note string) error
	EditNote(userID, numberNote int64, note string) error
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repositorer {
	return &repo{db: db}
}
