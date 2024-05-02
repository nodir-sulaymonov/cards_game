package models

import "time"

type AddRequest struct {
	AuthParams AuthData
}

type AuthData struct {
	Login    string
	Password string
}

type UserDate struct {
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate time.Time
	UpdateDate       time.Time
}