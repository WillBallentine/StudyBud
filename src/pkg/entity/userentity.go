package entity

type User struct {
	FirstName         string `bson:"first_name" json:"first_name"`
	LastName          string `bson:"last_name" json:"last_name"`
	Email             string `bson:"email" json:"email"`
	Password          string `bson:"pass" json:"pass"`
	School            string `bson:"school" json:"school"`
	SubscriptionLevel string `bson:"sub_level" json:"sub_level"`
}
