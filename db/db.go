package db

import (
	"cc-course-service/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// temporary dev db using sqlite
func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("course.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&model.Course{},
		&model.Subtopic{},
		&model.Content{},
		&model.Step{},
		&model.Feedback{},
	)

	DB = db
}
