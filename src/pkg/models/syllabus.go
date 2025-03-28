package model

type SyllabusData struct {
	CourseName  string       `json:"course_name"`
	CourseCode  string       `json:"course_code"`
	Instuctor   string       `json:"professor_name"`
	Assignments []Assignment `json:"assignments"`
	References  []Reference  `json:"required_texts"`
	School      string       `json:"school"`
	Semester    string       `json:"semester"`
}
