package handler

// tangkap parameter di handler
// handler ke service
// di service dia yang menentukan method repository mana yang di call
// repository : GetAll , Get By User ID
// db

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
