package rest_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/image"
	"github.com/solodba/iwechat/test/tools"
)

func TestCreateImage(t *testing.T) {
	imageReq := image.NewCreateImageRequest()
	imageReq.Model = "dall-e-3"
	imageReq.Prompt = "一个小狗"
	imageResp, err := client.CreateImage(ctx, imageReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(imageResp))
}
