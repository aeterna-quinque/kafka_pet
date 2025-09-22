package domain

import "kafka-pet/internal/models"

type User struct {
	Id   uint32
	Name string
	Age  uint8
}

func (u *User) ToUserModel() *models.User {
	return &models.User{
		Id:   u.Id,
		Name: u.Name,
		Age:  u.Age,
	}
}

func UserEntityFromUserModel(m *models.User) *User {
	u := User{
		Id:   m.Id,
		Name: m.Name,
		Age:  m.Age,
	}

	return &u
}
