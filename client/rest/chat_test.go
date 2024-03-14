package rest_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/iwechat/client/rest"
)

var (
	conf   *rest.Config
	client *rest.Client
	ctx    = context.Background()
)

func init() {
	conf = rest.NewConfig()
	client = rest.NewClient(conf)
}

func TestCreateChat(t *testing.T) {
	chatReq := chat.NewCreateChatRequest()
	chatReq.Model = "gpt-3.5-turbo"
	item1 := chat.NewMessagesItem()
	item1.Role = "system"
	item1.Content = "You are a helpful assistant."
	item2 := chat.NewMessagesItem()
	item2.Role = "user"
	item2.Content = "用中文介绍一下自己"
	chatReq.AddItems(item1, item2)
	chatResp, err := client.CreateChat(ctx, chatReq)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", chatResp)
	t.Log(chatResp.Data.Choices[0].Messages.Content)
}
