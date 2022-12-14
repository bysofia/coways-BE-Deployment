package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"email" validate:"required"`
}

type RegisterResponse struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"Email" form:"email" validate:"required"`
}
