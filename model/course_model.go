package model

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    string
	Prompt    string
	Title     string
	Desc      string
	Duration  string
	Theme     string
	Type      string
	IsDone    bool
	CreatedAt time.Time
	Subtopic  []Subtopic `gorm:"foreignKey:CourseID"`
}

type Subtopic struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Topic    string
	Desc     string
	CourseID uuid.UUID
	Course   Course
	IsDone   bool
	Content  Content `gorm:"foreignKey:SubID"`
}

type Content struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Opening string
	Closing string
	SubID   uuid.UUID
	Step    []Step `gorm:"foreignKey:ContentID"`
}

type Step struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Step      string
	ContentID uuid.UUID
}
