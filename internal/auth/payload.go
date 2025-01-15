package auth

type (
	RegisterResponse struct {
		Token string `json:"token"`
	}
	RegisterRequest struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
	}
)
