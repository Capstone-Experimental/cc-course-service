package repo

import (
	"cc-course-service/model"

	"gorm.io/gorm"
)

type CourseRepository struct {
	Db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		Db: db,
	}
}

func (repo *CourseRepository) GetAll() ([]model.Course, error) {
	var courses []model.Course
	result := repo.Db.Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (repo *CourseRepository) GetCourseByID(id string) (*map[string]interface{}, error) {
	var result = map[string]interface{}{}

	var course model.Course
	repo.Db.Where("id = ?", id).First(&course)

	result["id"] = course.ID
	result["title"] = course.Title
	result["desc"] = course.Desc
	result["duration"] = course.Duration

	var subtopics []model.Subtopic
	repo.Db.Where("course_id = ?", id).Find(&subtopics)

	var subMaps []map[string]interface{}

	for _, subtopic := range subtopics {
		var contents []model.Content
		repo.Db.Where("sub_id = ?", subtopic.ID).Find(&contents)

		var contentMaps []map[string]interface{}

		for _, content := range contents {
			var steps []model.Step
			repo.Db.Where("content_id = ?", content.ID).Find(&steps)

			// var stepMaps []map[string]interface{}
			var stepMaps []string

			for _, step := range steps {
				// var stepMap = map[string]interface{}{}
				// stepMap["id"] = step.ID
				// stepMap["step"] = step.Step
				var stepText = step.Step

				stepMaps = append(stepMaps, stepText)
			}

			var contentMap = map[string]interface{}{}
			// contentMap["id"] = content.ID
			contentMap["opening"] = content.Opening
			contentMap["steps"] = stepMaps
			contentMap["closing"] = content.Closing

			contentMaps = append(contentMaps, contentMap)
		}

		var subMap = map[string]interface{}{}
		subMap["id"] = subtopic.ID
		subMap["topic"] = subtopic.Topic
		subMap["contents"] = contentMaps

		subMaps = append(subMaps, subMap)
	}

	result["subtopics"] = subMaps

	return &result, nil
}

// marking subtopic as done
// if all subtopics are done, mark course as done
func (repo *CourseRepository) MarkSubtopicAsDone(id string) error {
	var subtopic model.Subtopic
	repo.Db.Where("id = ?", id).First(&subtopic)

	repo.Db.Model(&subtopic).Update("is_done", true)

	var course model.Course
	repo.Db.Where("id = ?", subtopic.CourseID).First(&course)

	var subtopics []model.Subtopic
	repo.Db.Where("course_id = ?", course.ID).Find(&subtopics)

	var isDone = true
	for _, subtopic := range subtopics {
		if !subtopic.IsDone {
			isDone = false
		}
	}

	if isDone {
		repo.Db.Model(&course).Update("is_done", true)
	}

	return nil
}
