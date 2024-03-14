package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/chat"
)

type CreateChatResp struct {
	Code int                      `json:"code"`
	Data *chat.CreateChatResponse `json:"data"`
}

func NewCreateChatResp() *CreateChatResp {
	return &CreateChatResp{
		Data: chat.NewCreateChatResponse(),
	}
}

func (c *Client) CreateChat(ctx context.Context, req *chat.CreateChatRequest) (*CreateChatResp, error) {
	createChatResponse := NewCreateChatResp()
	err := c.c.Post("chat").Body(req).Do(ctx).Into(createChatResponse)
	if err != nil {
		return nil, err
	}
	return createChatResponse, nil
}
