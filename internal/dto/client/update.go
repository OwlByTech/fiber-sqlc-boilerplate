package dto

type UpdateClientByIdReq struct {
	ClientId  int64  `json:"id"`
	GivenName string `json:"givenName"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type UpdateClientByIdRes struct {
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
