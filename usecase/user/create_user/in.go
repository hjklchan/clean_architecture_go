package create_user

import "time"

type Input struct {
	Name        string
	Email       string
	DateOfBirth time.Time
	Password    string
}
