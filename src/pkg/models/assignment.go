package model

type Assignment struct {
	Title       string
	Description string
	DueDate     string
	Refs        []Reference
	Type        string
}
