package dto

import "kafka-pet/internal/domain"

type CreateUserRequest struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (r *CreateUserRequest) ToUser() *domain.User {
	return &domain.User{
		Id:   r.Id,
		Name: r.Name,
		Age:  r.Age,
	}
}
