package model

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId    string    `json:"user_id"`
	CourseId  string    `json:"course_id"`
	Prompt    string    `json:"prompt"`
	Feedback  string    `json:"feedback"`
	Rating    float32   `json:"rating" gorm:"type:numeric(2,1)"`
	CreatedAt time.Time `json:"created_at"`
}
