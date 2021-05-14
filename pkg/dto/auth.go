package dto

// LoginDTO defined the /login payload
type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"password,min=6,max=34"`
	Remember bool   `json:"remember"`
}

// RegisterDTO defined the /register payload
type RegisterDTO struct {
	Username  string `json:"username" validate:"required,min=3,max=21"`
	Password  string `json:"password" validate:"required,min=6,max=34"`
	Email     string `json:"email" validate:"required,email"`
	Agreement bool   `json:"agreement" validate:"required"`
}

// AuthResponse
type AuthResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
