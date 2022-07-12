package dto

type UserLoginDTO struct {
	Email    string
	Password string
}

type UserUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	Firstname string `json:"firstname" form:"name" binding:"required"`
	Lastname  string `json:"lastname" form:"name" binding:"required"`
	Email     string `json:"email" from:"email" binding:"required,email"`
	Password  string `json:"password" from:"password" binding:"required"`
}
