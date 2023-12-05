package model

type CourseRaw struct {
	Course string `json:"course"`
	Prompt string `json:"prompt"`
}

type FeedbackRaw struct {
	Feedback string  `json:"feedback"`
	Rating   float32 `json:"rating"`
}
