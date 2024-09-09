package dto

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordTokenReq struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type ValidateResetPasswordTokenReq struct {
	Token string `json:"token" validate:"required"`
}
