package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	FirstName         string                `bson:"first_name" json:"first_name"`
	LastName          string                `bson:"last_name" json:"last_name"`
	Email             string                `bson:"email" json:"email"`
	Password          string                `bson:"pass" json:"pass"`
	SessionToken      string                `bson:"session_token" json:"session_token"`
	CSRFToken         string                `bson:"csrf_token" json:"csrf_token"`
	School            string                `bson:"school" json:"school"`
	SubscriptionLevel string                `bson:"sub_level" json:"sub_level"`
	Cohorts           []primitive.ObjectID  `bson:"cohorts" json:"cohorts"`
	Syllabi           []primitive.ObjectID  `bson:"syllabi" json:"syllabi"`
	StudyPlans        []primitive.ObjectID  `bson:"study_plans" json:"study_plans"`
	UserPreferences   UserPreferencesEntity `bson:"user_preferences" json:"user_preferences"`
}

type UserPreferencesEntity struct {
	ScheduleRules      ScheduleRulesEntity `bson:"schedule_rules" json:"schedule_rules"`
	EmailNotifications bool                `bson:"email_notifications" json:"email_notifications"`
	TextNotifications  bool                `bson:"text_notifications" json:"text_notifications"`
	DefaultClass       primitive.ObjectID  `bson:"default_class" json:"default_class"`
}

type ScheduleRulesEntity struct {
	PreferredStudyTimes []StudyTimeRangeEntity `bson:"preferred_study_times"`
	MaxDailyStudyHours  int                    `bson:"max_daily_study_hours" json:"max_daily_study_hours,omitempty"`
	BreakDuration       string                 `bson:"break_duration" json:"break_duration,omitempty"`
}

type StudyTimeRangeEntity struct {
	DayOfWeek string `bson:"day_of_week" json:"day_of_week"`
	StartTime string `bson:"start_time" json:"start_time"`
	EndTime   string `bson:"end_time" json:"end_time"`
}
