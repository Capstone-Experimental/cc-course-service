package model

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    string     `json:"user_id"`
	Prompt    string     `json:"prompt"`
	Title     string     `json:"title"`
	Desc      string     `json:"desc"`
	Duration  string     `json:"duration"`
	Theme     string     `json:"theme"`
	Type      string     `json:"type"`
	Progress  float32    `json:"progress"`
	IsDone    bool       `json:"is_done"`
	CreatedAt time.Time  `json:"created_at"`
	Subtopic  []Subtopic `gorm:"foreignKey:CourseID" json:"subtitle"`
}

type Subtopic struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"sub_id"`
	Topic    string    `json:"topic"`
	Desc     string    `json:"shortdesc"`
	CourseID uuid.UUID `json:"course_id"`
	Course   Course    `gorm:"foreignKey:CourseID"`
	IsDone   bool      `json:"is_done"`
	Content  Content   `gorm:"foreignKey:SubID"`
}

type Content struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"content_id"`
	Opening string    `json:"opening"`
	Closing string    `json:"closing"`
	SubID   uuid.UUID `json:"sub_id"`
	Step    []Step    `gorm:"foreignKey:ContentID" json:"steps"`
}

type Step struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Step      string    `json:"step"`
	ContentID uuid.UUID `json:"content_id"`
}
