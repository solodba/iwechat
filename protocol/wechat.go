package protocol

import (
	"context"

	"github.com/solodba/iwechat/apps/wechat"
	"github.com/solodba/mcube/apps"
)

var (
	ctx = context.Background()
)

// Wechat服务结构体
type WechatService struct {
	svc wechat.Service
}

// Wechat服务结构体初始化
func NewWechatService() *WechatService {
	return &WechatService{
		svc: apps.GetInternalApp(wechat.AppName).(wechat.Service),
	}
}

// Wechat服务启动方法
func (w *WechatService) Start() error {
	return w.svc.ChatBot(ctx)
}

// Wechat服务停止方法
func (s *WechatService) Stop() error {
	return nil
}
