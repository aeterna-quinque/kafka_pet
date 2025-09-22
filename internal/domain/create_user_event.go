package domain

import "time"

type CreateUserEvent struct {
	UserId    uint32    `json:"user_id"`
	Name      string    `json:"name"`
	Age       uint8     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}
