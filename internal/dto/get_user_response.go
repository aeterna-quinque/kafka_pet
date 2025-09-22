package dto

import "kafka-pet/internal/domain"

type GetUserResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func GetUserResponseFromUser(u *domain.User) *GetUserResponse {
	return &GetUserResponse{
		Id:   u.Id,
		Name: u.Name,
		Age:  u.Age,
	}
}
