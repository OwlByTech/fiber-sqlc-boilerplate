package dto

type LoginClientReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginClientRes struct {
	Token string `json:"token"`
}
