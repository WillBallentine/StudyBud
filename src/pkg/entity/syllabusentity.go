package entity

type SyllabusDataEntity struct {
	CourseName    string             `bson:"course_name,omitempty" json:"course_name"`
	CourseCode    string             `bson:"course_code" json:"course_code"`
	Instructor    string             `bson:"professor_name" json:"professor_name"`
	Assignments   []AssignmentEntity `bson:"assignments" json:"assignments"`
	RequiredTexts []ReferenceEntity  `bson:"required_texts" json:"required_texts"`
	School        string             `bson:"school,omitempty" json:"school"`
	Semester      string             `bson:"semester,omitempty" json:"semester"`
}

type AssignmentEntity struct {
	Title       string            `bson:"title" json:"title"`
	Description string            `bson:"desc" json:"desc"`
	DueDate     string            `bson:"due_date,omitempty" json:"due_date"`
	Refs        []ReferenceEntity `bson:"related_texts,omitempty" json:"related_texts"`
}

type ReferenceEntity struct {
	Title  string `bson:"title" json:"title"`
	Author string `bson:"author" json:"author"`
	Link   string `bson:"link,omitempty" json:"link"`
}
