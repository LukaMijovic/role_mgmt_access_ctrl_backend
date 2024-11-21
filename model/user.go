package model

import "time"

type User struct {
	UserID               int64
	Firstname            string    `binding:"required"`
	Lastname             string    `binding:"required"`
	Email                string    `binding:"required"`
	Telephone            string    `binding:"required"`
	Birthdate            time.Time `binding:"required"`
	UserRegistrationDate time.Time
	roleID               int64
}

func NewUser(firstname string, lastname string, email string, telephone string, birthdate time.Time, registrationDate time.Time) *User {
	return &User{
		Firstname:            firstname,
		Lastname:             lastname,
		Email:                email,
		Telephone:            telephone,
		Birthdate:            birthdate,
		UserRegistrationDate: registrationDate,
	}
}

func (u *User) GetID() int64 {
	return u.UserID
}

func (u *User) SetID(userID int64) {
	u.UserID = userID
}

func (u *User) GetUserRegistrationDate() time.Time {
	return u.UserRegistrationDate
}

func (u *User) SetUserRegistraionDate() {
	u.UserRegistrationDate = time.Now()
}

func (u *User) GetRoleID() int64 {
	return u.roleID
}

func (u *User) SetRoleID(roleID int64) {
	u.roleID = roleID
}
