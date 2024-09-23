package dto

type UserAccessDTO struct {
	UserId int64  `binding:required`
	IMEI   string `binding:required`
}
