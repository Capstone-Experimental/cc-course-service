package model

type Dashboard struct {
	Course           []Course `json:"courses"`
	Progress         []Course `json:"progress"`
	CourseCompleted  int      `json:"course_completed"`
	CourseInProgress int      `json:"course_progress"`
}
