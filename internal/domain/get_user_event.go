package domain

import "time"

type GetUserEvent struct {
	UserId      uint32    `json:"user_id"`
	RequestedAt time.Time `json:"requested_at"`
}
