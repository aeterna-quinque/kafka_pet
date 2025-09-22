package dto

import "kafka-pet/internal/domain"

type GetUserResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (r *GetUserResponse) FromUser(u *domain.User) {
	r.Id = u.Id
	r.Name = u.Name
	r.Age = u.Age
}
