package repo

import (
	"cc-course-service/model"

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

// get all feedbacks
func (repo *FeedbackRepository) GetAllFeedbacks(offset, pageSize int) ([]model.Feedback, error) {

	if offset > 0 && pageSize > 0 {
		// Query for the current page of items
		rows, err := repo.Db.Raw("SELECT * FROM feedbacks LIMIT ? OFFSET ?", pageSize, offset).Rows()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var feedbacks []model.Feedback
		for rows.Next() {
			var feedback model.Feedback
			repo.Db.ScanRows(rows, &feedback)
			feedbacks = append(feedbacks, feedback)
		}

		return feedbacks, nil
	}

	var feedbacks []model.Feedback
	result := repo.Db.Find(&feedbacks)

	if result.Error != nil {
		return nil, result.Error
	}
	return feedbacks, nil
}
