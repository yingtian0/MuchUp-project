package domain

import "time"

type User struct {
	UserID    string
	Name      string
	Profile   Profile
	CreatedAt time.Time
}

type Profile struct {
	Avatar string
	Bio    string
}
