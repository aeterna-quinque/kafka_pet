package dto

import "kafka-pet/internal/domain"

type GetMessagesResponse struct {
	Messages []string `json:"messages"`
}

func MessagesToGetMessagesResponse(s *domain.Messages) *GetMessagesResponse {
	resp := &GetMessagesResponse{
		Messages: make([]string, len(s.Messages)),
	}

	copy(resp.Messages, s.Messages)

	return resp
}
