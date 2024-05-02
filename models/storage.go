package models

import "time"

type User struct {
	Id               int
	Name             string
	Surname          string
	Login            string
	HashPassword     string
	Status           string
	Role             string
	RegistrationDate time.Time
	UpdateDate       time.Time
}