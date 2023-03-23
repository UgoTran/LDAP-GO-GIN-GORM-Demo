package entity

import "time"

type RpUser struct {
	Username  string
	Avatar    string
	Email     string
	FullName  string
	UpdatedAt time.Time
	Status    bool
	Source    string
	Password  string
}
