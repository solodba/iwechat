package rest

import (
	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/ichatgpt/apps/image"
)

type CreateChatResp struct {
	Code int                      `json:"code"`
	Data *chat.CreateChatResponse `json:"data"`
}

type CreateImageResp struct {
	Code int                        `json:"code"`
	Data *image.CreateImageResponse `json:"data"`
}

func NewCreateChatResp() *CreateChatResp {
	return &CreateChatResp{
		Data: chat.NewCreateChatResponse(),
	}
}

func NewCreateImageResp() *CreateImageResp {
	return &CreateImageResp{
		Data: image.NewCreateImageResponse(),
	}
}
