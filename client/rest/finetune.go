package rest

import (
	"context"

	"github.com/solodba/ichatgpt/apps/finetune"
)

func (c *Client) CreateFineTuneJob(ctx context.Context, req *finetune.CreateFineTuneJobRequest) (*CreateFineTuneJobResp, error) {
	createFineTuneJobResp := NewCreateFineTuneJobResp()
	err := c.c.Post("finetune").Body(req).Do(ctx).Into(createFineTuneJobResp)
	if err != nil {
		return nil, err
	}
	return createFineTuneJobResp, nil
}
