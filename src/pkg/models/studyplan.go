package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudyPlan struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CourseName    string             `bson:"course_name,omitempty" json:"course_name"`
	CourseCode    string             `bson:"course_code" json:"course_code"`
	Semester      string             `bson:"semester,omitempty" json:"semester"`
	StudyBlocks   []StudyBlock       `bson:"study_blocks" json:"study_blocks"`
	GeneratedAt   string             `bson:"generated_at" json:"generated_at"`
	ScheduleRules ScheduleRules      `bson:"schedule_rules,omitempty" json:"schedule_rules"`
}

type StudyBlock struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title         string             `bson:"title" json:"title"`
	Description   string             `bson:"description,omitempty" json:"description"`
	StartDate     string             `bson:"start_date" json:"start_date"`
	DueDate       string             `bson:"due_date,omitempty" json:"due_date"`
	Assignment    *Assignment        `bson:"assignment,omitempty" json:"assignment"`
	Refs          []Reference        `bson:"related_texts,omitempty" json:"related_texts"`
	Completed     bool               `bson:"completed" json:"completed"`
	Priority      int                `bson:"priority" json:"priority"`
	EstimatedTime string             `bson:"estimated_time" json:"estimated_time"`
	Recurring     bool               `bson:"recurring,omitempty" json:"recurring"`
	Recurrence    *RecurrenceRule    `bson:"recurrence,omitempty" json:"recurrence"`
}

type RecurrenceRule struct {
	Frequency string   `bson:"frequency" json:"frequency"`
	Interval  int      `bson:"interval" json:"interval"`
	Weekdays  []string `bson:"weekdays,omitempty" json:"weekdays,omitempty"`
	EndDate   string   `bson:"end_date,omitempty" json:"end_date,omitempty"`
}

type ScheduleRules struct {
	PreferredStudyTimes []StudyTimeRange `bson:"preferred_study_times,omitempty"`
	MaxDailyStudyHours  int              `bson:"max_daily_study_hours,omitempty" json:"max_daily_study_hours,omitempty"`
	BreakDuration       string           `bson:"break_duration,omitempty" json:"break_duration,omitempty"`
}

type StudyTimeRange struct {
	DayOfWeek string `bson:"day_of_week" json:"day_of_week"`
	StartTime string `bson:"start_time" json:"start_time"`
	EndTime   string `bson:"end_time" json:"end_time"`
}
