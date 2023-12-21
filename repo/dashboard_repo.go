package repo

import (
	"cc-course-service/model"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	Db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{
		Db: db,
	}
}

// get dashboard according to user id
func (repo *DashboardRepository) GetDashboard(uid string) (*model.Dashboard, error) {
	var dashboard model.Dashboard

	var courses []model.Course
	result := repo.Db.Where("user_id = ?", uid).Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	var progress []model.Course
	result = repo.Db.Raw("SELECT * FROM courses WHERE user_id = ? AND progress != 0 AND progress != 100", uid).Find(&progress)
	if result.Error != nil {
		return nil, result.Error
	}

	var completed []model.Course
	result = repo.Db.Raw("SELECT * FROM courses WHERE user_id = ? AND progress = 100", uid).Find(&completed)
	if result.Error != nil {
		return nil, result.Error
	}

	dashboard.Course = courses
	dashboard.Progress = progress
	dashboard.CourseCompleted = len(completed)
	dashboard.CourseInProgress = len(courses) - len(completed)

	return &dashboard, nil
}
