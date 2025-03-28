package model

type Assignment struct {
	Title       string      `json:"title"`
	Description string      `json:"desc"`
	DueDate     string      `json:"due_date"`
	Refs        []Reference `json:"related_texts"`
}
