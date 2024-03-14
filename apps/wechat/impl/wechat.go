package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/ichatgpt/apps/image"
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

func (i *impl) ChatBot(ctx context.Context) error {
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
		flag := IsInRemarkNameList(sendUser.RemarkName, i.c.WeChat.RemarkNameList)
		if flag && msg.IsText() {
			contentSegList := strings.Split(msg.Content, "-")
			chatgptClient := rest.NewClient(rest.NewConfig())
			if contentSegList[0] == "image" {
				imageReq := image.NewCreateImageRequest()
				imageReq.Model = "dall-e-3"
				imageReq.Prompt = contentSegList[1]
				ctx, cancel := context.WithTimeout(context.Background(), time.Hour*1)
				defer cancel()
				imageResp, err := chatgptClient.CreateImage(ctx, imageReq)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				if len(imageResp.Data.Data) == 0 {
					msg.ReplyText("内容可能违法,请重新提问,谢谢!")
					return
				}
				httpClient := http.Client{
					Timeout: time.Duration(10 * time.Minute),
				}
				httpResp, err := httpClient.Get(imageResp.Data.Data[0].Url)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				msg.ReplyImage(httpResp.Body)
				return
			} else {
				chatReq := chat.NewCreateChatRequest()
				chatReq.Model = "gpt-3.5-turbo"
				item1 := chat.NewMessagesItem()
				item1.Role = "system"
				item1.Content = "You are a helpful assistant."
				item2 := chat.NewMessagesItem()
				item2.Role = "user"
				item2.Content = msg.Content
				chatReq.AddItems(item1, item2)
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
				defer cancel()
				chatResp, err := chatgptClient.CreateChat(ctx, chatReq)
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

func IsInRemarkNameList(name string, nameList []string) bool {
	for _, item := range nameList {
		if item == name {
			return true
		}
	}
	return false
}
