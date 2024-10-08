package repository

import "github.com/BelyaevEI/platform_common/pkg/db"

// Имплементация репо слоя
type Repositorer interface {
	CreateNote(userID int64, note string) error
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repositorer {
	return &repo{db: db}
}
