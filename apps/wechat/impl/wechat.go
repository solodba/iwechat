package impl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/solodba/ichatgpt/apps/audio"
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
		contentSegList := strings.Split(msg.Content, "-")
		chatgptClient := rest.NewClient(rest.NewConfig())
		if flag && msg.IsText() {
			switch contentSegList[0] {
			case "图片":
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
			case "文转音":
				speechReq := audio.NewCreateSpeechRequest()
				speechReq.Model = "tts-1-hd"
				speechReq.Input = contentSegList[1]
				speechReq.Voice = "alloy"
				speechReq.FilePath = "audio"
				speechReq.FileName = "voice.mp3"
				ctx, cancel := context.WithTimeout(context.Background(), time.Hour*1)
				defer cancel()
				_, err := chatgptClient.CreateSpeech(ctx, speechReq)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				voiceFilePath := `C:\Users\Admin\Desktop\ichatgpt\audio\voice.mp3`
				f, err := os.Open(voiceFilePath)
				defer f.Close()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				msg.ReplyFile(f)
			default:
				chatReq := chat.NewCreateChatRequest()
				chatReq.Model = "gpt-4-0125-preview"
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
		if flag && msg.IsVoice() {
			voiceFilePath := `C:\Users\Admin\Desktop\ichatgpt\audio\voice.mp3`
			voiceResp, err := msg.GetVoice()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			f, err := os.OpenFile(voiceFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer f.Close()
			_, err = io.Copy(f, voiceResp.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			// transcriptionReq := audio.NewCreateTranscriptionRequest()
			// transcriptionReq.Model = "whisper-1"
			// transcriptionReq.Language = "zh"
			// transcriptionReq.ResponseFormat = "json"
			// transcriptionReq.FilePath = "audio"
			// transcriptionReq.FileName = "voice.mp3"
			// ctx, cancel := context.WithTimeout(context.Background(), time.Hour*1)
			// defer cancel()
			// transcriptionResp, err := chatgptClient.CreateTranscription(ctx, transcriptionReq)
			// if err != nil {
			// 	fmt.Println(err.Error())
			// 	return
			// }
			// msg.ReplyText(transcriptionResp.Data.Text)
			translationReq := audio.NewCreateTranslationRequest()
			translationReq.Model = "whisper-1"
			translationReq.FileName = "voice.mp3"
			translationReq.FilePath = "audio"
			translationReq.ResponseFormat = "json"
			translationReq.Temperature = 0.2
			ctx, cancel := context.WithTimeout(context.Background(), time.Hour*1)
			defer cancel()
			translationResp, err := chatgptClient.CreateTranslation(ctx, translationReq)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			msg.ReplyText(translationResp.Data.Text)
			return
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
