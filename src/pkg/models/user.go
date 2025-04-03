package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID   `json:"id"`
	FirstName         string               `json:"first_name"`
	LastName          string               `json:"last_name"`
	Email             string               `json:"email"`
	Password          string               `json:"pass"`
	School            string               `json:"school"`
	SubscriptionLevel string               `json:"sub_level"`
	Cohorts           []primitive.ObjectID `json:"cohorts"`
	Syllabi           []primitive.ObjectID `json:"syllabi"`
	StudyPlans        []primitive.ObjectID `json:"study_plans"`
	UserPreferences   UserPreferences      `json:"user_preferences"`
}

type UserPreferences struct {
	ScheduleRules      ScheduleRules      `bson:"schedule_rules" json:"schedule_rules"`
	EmailNotifications bool               `bson:"email_notifications" json:"email_notifications"`
	TextNotifications  bool               `bson:"text_notifications" json:"text_notifications"`
	DefaultClass       primitive.ObjectID `bson:"default_class" json:"default_class"`
}
