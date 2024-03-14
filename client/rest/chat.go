package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/chat"
)

func (c *Client) CreateChat(ctx context.Context, req *chat.CreateChatRequest) (*CreateChatResp, error) {
	createChatResponse := NewCreateChatResp()
	err := c.c.Post("chat").Body(req).Do(ctx).Into(createChatResponse)
	if err != nil {
		return nil, err
	}
	return createChatResponse, nil
}
