package dto

type CreateClientReq struct {
	Email     string `json:"email" validate:"required,email"`
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}

type CreateClientRes struct {
	Token string `json:"token"`
}
