package rest

import (
	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/ichatgpt/apps/file"
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

type CreateSpeechResp struct {
	Code int                         `json:"code"`
	Data *audio.CreateSpeechResponse `json:"data"`
}

type CreateTranscriptionResp struct {
	Code int                        `json:"code"`
	Data *audio.CreateAudioResponse `json:"data"`
}

type CreateTranslationResp struct {
	Code int                        `json:"code"`
	Data *audio.CreateAudioResponse `json:"data"`
}

type UploadFileResp struct {
	Code int                    `json:"code"`
	Data *file.FileResponseItem `json:"data"`
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

func NewCreateSpeechResp() *CreateSpeechResp {
	return &CreateSpeechResp{}
}

func NewCreateTranscriptionResp() *CreateTranscriptionResp {
	return &CreateTranscriptionResp{}
}

func NewCreateTranslationResp() *CreateTranslationResp {
	return &CreateTranslationResp{}
}

func NewUploadFileResp() *UploadFileResp {
	return &UploadFileResp{}
}
