package dto

type LoginDTO struct {
	Username string `json:"username" from:"username" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}
