package dtos

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,alphaunicode,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,numeric,min=6,max=20"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"alphaunicode,min=2"`
	Email    string `json:"email" binding:"email,min=4"`
	Phone    string `json:"phone" binding:"numeric,min=6,max=20"`
	Password string `json:"password"`
}