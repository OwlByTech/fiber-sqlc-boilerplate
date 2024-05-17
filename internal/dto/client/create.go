package dto

type CreateClientReq struct {
	Email     string `json:"email" validate:"required,email"`
	GivenName string `json:"givenName" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
	Surname   string `json:"surname" validate:"required"`
}

type CreateClientRes struct {
	Id        int64
	Email     string
	GivenName string
	Surname   string
}
