package dtos

import "time"

type UserUpdateDto struct {
	ID           uint
	NameSurname  string
	Email        string
	PasswordHash string
	UpdatedOn    *time.Time
	UpdatedBy    string
}
