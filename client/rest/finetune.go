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

func (c *Client) RetrieveFineTuneJob(ctx context.Context, req *finetune.RetrieveFineTuneJobRequest) (*RetrieveFineTuneJobResp, error) {
	retrieveFineTuneJobResp := NewRetrieveFineTuneJobResp()
	err := c.c.Get("finetune/" + req.FineTuningJobId).Body(req).Do(ctx).Into(retrieveFineTuneJobResp)
	if err != nil {
		return nil, err
	}
	return retrieveFineTuneJobResp, nil
}
