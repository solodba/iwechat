package rest

import "github.com/solodba/ichatgpt/apps/chat"

type CreateChatResp struct {
	Code int                      `json:"code"`
	Data *chat.CreateChatResponse `json:"data"`
}

func NewCreateChatResp() *CreateChatResp {
	return &CreateChatResp{
		Data: chat.NewCreateChatResponse(),
	}
}
