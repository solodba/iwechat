package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/file"
)

func (c *Client) UploadFile(ctx context.Context, req *file.UploadFileRequest) (*UploadFileResp, error) {
	uploadFileResponse := NewUploadFileResp()
	err := c.c.Post("file").Body(req).Do(ctx).Into(uploadFileResponse)
	if err != nil {
		return nil, err
	}
	return uploadFileResponse, nil
}
