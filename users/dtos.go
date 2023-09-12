package users

type CreateUserDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserDTO CreateUserDTO
