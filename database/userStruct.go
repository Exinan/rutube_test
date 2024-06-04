package database

import (
	"time"
)

type UserDataForBirth struct {
	Name      string
	Surname   string
	BirthDate time.Time
}
