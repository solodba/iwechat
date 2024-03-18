package rest_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/finetune"
	"github.com/solodba/iwechat/test/tools"
)

func TestCreateFineTuneJob(t *testing.T) {
	fineTuneJobReq := finetune.NewCreateFineTuneJobRequest()
	fineTuneJobReq.Model = "gpt-3.5-turbo-0125"
	fineTuneJobReq.TrainingFile = "file-Z4Mgod4VdspC1UCBbJtC2Nhw"
	fineTuneJobResp, err := client.CreateFineTuneJob(ctx, fineTuneJobReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(fineTuneJobResp))
}
