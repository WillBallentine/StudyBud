package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FirstName         string               `bson:"first_name" json:"first_name"`
	LastName          string               `bson:"last_name" json:"last_name"`
	Email             string               `bson:"email" json:"email"`
	Password          string               `bson:"pass" json:"pass"`
	SessionToken      string               `bson:"session_token" json:"session_token"`
	CSRFToken         string               `bson:"csrf_token" json:"csrf_token"`
	School            string               `bson:"school" json:"school"`
	SubscriptionLevel string               `bson:"sub_level" json:"sub_level"`
	Cohorts           []primitive.ObjectID `bson:"cohorts" json:"cohorts"`
	Syllabi           []primitive.ObjectID `bson:"syllabi" json:"syllabi"`
	StudyPlans        []primitive.ObjectID `bson:"study_plans" json:"study_plans"`
}
