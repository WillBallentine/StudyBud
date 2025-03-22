package model

type User struct {
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	Email             string   `json:"email"`
	Password          string   `json:"pass"`
	School            string   `json:"school"`
	SubscriptionLevel string   `json:"sub_level"`
	Cohorts           []string `json:"cohorts"`
}
