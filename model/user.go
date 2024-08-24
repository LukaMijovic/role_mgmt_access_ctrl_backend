package model

import "time"

type User struct {
	userID              int64
	firstname           string
	lastname            string
	email               string
	telephone           string
	birthdate           time.Time
	userRegistraionDate time.Time
	roleID              int64
}

func (u *User) GetID() int64 {
	return u.userID
}

func (u *User) GetFirstname() string {
	return u.firstname
}

func (u *User) GetLastname() string {
	return u.lastname
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetTelephone() string {
	return u.telephone
}

func (u *User) GetBirthdate() time.Time {
	return u.birthdate
}

func (u *User) GetUserRegistrationDate() time.Time {
	return u.userRegistraionDate
}

func (u *User) GetRoleID() int64 {
	return u.roleID
}
