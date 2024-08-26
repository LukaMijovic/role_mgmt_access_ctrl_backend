package dto

type AdminCredentialsDTO struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
