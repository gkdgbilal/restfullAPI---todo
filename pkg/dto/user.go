package dto

type UserDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email,max=254"`
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
