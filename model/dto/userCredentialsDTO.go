package dto

type UserCredentialsDTO struct {
	User_ID  int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
