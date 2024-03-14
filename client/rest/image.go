package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/image"
)

func (c *Client) CreateImage(ctx context.Context, req *image.CreateImageRequest) (*CreateImageResp, error) {
	createImageResponse := NewCreateImageResp()
	err := c.c.Post("image").Body(req).Do(ctx).Into(createImageResponse)
	if err != nil {
		return nil, err
	}
	return createImageResponse, nil
}
