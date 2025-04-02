package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SyllabusData struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	CourseName  string             `json:"course_name"`
	CourseCode  string             `json:"course_code"`
	Instuctor   string             `json:"professor_name"`
	Assignments []Assignment       `json:"assignments"`
	References  []Reference        `json:"required_texts"`
	School      string             `json:"school"`
	Semester    string             `json:"semester"`
}
