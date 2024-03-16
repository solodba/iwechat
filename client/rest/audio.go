package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/audio"
)

func (c *Client) CreateSpeech(ctx context.Context, req *audio.CreateSpeechRequest) (*CreateSpeechResp, error) {
	createAudioResponse := NewCreateSpeechResp()
	err := c.c.Post("audio/speech").Body(req).Do(ctx).Into(createAudioResponse)
	if err != nil {
		return nil, err
	}
	return createAudioResponse, nil
}

func (c *Client) CreateTranscription(ctx context.Context, req *audio.CreateTranscriptionRequest) (*CreateTranscriptionResp, error) {
	createTranscriptionResponse := NewCreateTranscriptionResp()
	err := c.c.Post("audio/transcription").Body(req).Do(ctx).Into(createTranscriptionResponse)
	if err != nil {
		return nil, err
	}
	return createTranscriptionResponse, nil
}
