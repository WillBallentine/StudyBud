package entity

type User struct {
	ID                string `bson:"_id" json:"_id"`
	FirstName         string `bson:"first_name" json:"first_name"`
	LastName          string `bson:"last_name" json:"last_name"`
	Email             string `bson:"email" json:"email"`
	Password          string `bson:"pass" json:"pass"`
	SessionToken      string `bson:"session_token" json:"session_token"`
	CSRFToken         string `bson:"csrf_token" json:"csrf_token"`
	School            string `bson:"school" json:"school"`
	SubscriptionLevel string `bson:"sub_level" json:"sub_level"`
}
