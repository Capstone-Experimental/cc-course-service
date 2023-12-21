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

// get all courses
func (repo *CourseRepository) GetAll(offset, pageSize int) ([]model.Course, error) {

	if offset > 0 && pageSize > 0 {
		// Query for the current page of items
		rows, err := repo.Db.Raw("SELECT * FROM courses LIMIT ? OFFSET ?", pageSize, offset).Rows()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var courses []model.Course
		for rows.Next() {
			var course model.Course
			repo.Db.ScanRows(rows, &course)
			courses = append(courses, course)
		}

		return courses, nil
	}

	var courses []model.Course
	result := repo.Db.Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

// get all courses according to user id
func (repo *CourseRepository) MyCourse(uid string, offset, pageSize int) ([]model.Course, error) {

	if offset > 0 && pageSize > 0 {
		// Query for the current page of items
		rows, err := repo.Db.Raw("SELECT * FROM courses WHERE user_id = ? LIMIT ? OFFSET ?", uid, pageSize, offset).Rows()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var courses []model.Course
		for rows.Next() {
			var course model.Course
			repo.Db.ScanRows(rows, &course)
			courses = append(courses, course)
		}

		return courses, nil
	}

	var courses []model.Course
	result := repo.Db.Where("user_id = ?", uid).Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

// get course by id
func (repo *CourseRepository) GetCourseByID(id string) (*map[string]interface{}, error) {
	var result = map[string]interface{}{}

	var course model.Course
	repo.Db.Where("id = ?", id).First(&course)

	result["id"] = course.ID
	result["title"] = course.Title
	result["desc"] = course.Desc
	result["type"] = course.Type
	result["theme"] = course.Theme
	result["duration"] = course.Duration
	result["is_done"] = course.IsDone

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

			var stepMaps []string

			for _, step := range steps {
				var stepText = step.Step

				stepMaps = append(stepMaps, stepText)
			}

			var contentMap = map[string]interface{}{}
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

	var subtopicMarkedAsDone float32 = 0

	var isDone = true
	for _, subtopic := range subtopics {
		if !subtopic.IsDone {
			isDone = false
		} else {
			subtopicMarkedAsDone++
		}
	}

	//auto update progress according to subtopic marked as done (done subtopic / total subtopic)
	var progress = subtopicMarkedAsDone / float32(len(subtopics)) * 100

	repo.Db.Model(&course).Update("progress", progress)

	if isDone {
		repo.Db.Model(&course).Update("is_done", true)
	}

	return nil
}
