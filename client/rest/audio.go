package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/audio"
)

func (c *Client) CreateSpeech(ctx context.Context, req *audio.CreateSpeechRequest) (*CreateSpeechResp, error) {
	createSpeechResponse := NewCreateSpeechResp()
	err := c.c.Post("audio/speech").Body(req).Do(ctx).Into(createSpeechResponse)
	if err != nil {
		return nil, err
	}
	return createSpeechResponse, nil
}
