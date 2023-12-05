package repo

import (
	"gorm.io/gorm"
)

type FeedbackRepository struct {
	Db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{
		Db: db,
	}
}
