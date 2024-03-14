package impl

import (
	"context"
	"fmt"

	"github.com/eatmoreapple/openwechat"
	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/iwechat/client/rest"
)

func (i *impl) WechatLogin(ctx context.Context) error {
	i.bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	if err := i.bot.Login(); err != nil {
		return err
	}
	return nil
}

func (i *impl) GetWechatUsers(ctx context.Context) (openwechat.Friends, error) {
	self, err := i.bot.GetCurrentUser()
	if err != nil {
		return nil, err
	}
	return self.Friends()
}

func (i *impl) GetWechatGroups(ctx context.Context) (openwechat.Groups, error) {
	self, err := i.bot.GetCurrentUser()
	if err != nil {
		return nil, err
	}
	return self.Groups()
}

func (i *impl) ChatTextBot(ctx context.Context) error {
	err := i.WechatLogin(ctx)
	if err != nil {
		return err
	}
	_, err = i.GetWechatUsers(ctx)
	if err != nil {
		return err
	}
	i.bot.MessageHandler = func(msg *openwechat.Message) {
		sendUser, err := msg.Sender()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		chatgptClient := rest.NewClient(rest.NewConfig())
		for _, remarkname := range i.c.WeChat.RemarkNameList {
			if msg.IsText() && sendUser.RemarkName == remarkname {
				chatReq := chat.NewCreateChatRequest()
				chatReq.Model = "gpt-3.5-turbo"
				item1 := chat.NewMessagesItem()
				item1.Role = "system"
				item1.Content = "You are a helpful assistant."
				item2 := chat.NewMessagesItem()
				item2.Role = "user"
				item2.Content = msg.Content
				chatReq.AddItems(item1, item2)
				chatResp, err := chatgptClient.CreateChat(context.Background(), chatReq)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				msg.ReplyText(chatResp.Data.Choices[0].Messages.Content)
				return
			}
		}
	}
	i.bot.Block()
	return nil
}
