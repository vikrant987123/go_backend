package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB string `json:"dob" validate:"required"`
}