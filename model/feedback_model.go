package model

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId    string
	CourseId  string
	Feedback  string
	Rating    float32
	CreatedAt time.Time
}
