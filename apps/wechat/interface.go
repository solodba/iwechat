package wechat

import (
	"context"

	"github.com/eatmoreapple/openwechat"
)

// 模块名称
const (
	AppName = "wechat"
)

// wechat功能接口
type Service interface {
	WechatLogin(context.Context) error
	GetWechatUsers(context.Context) (openwechat.Friends, error)
	GetWechatGroups(context.Context) (openwechat.Groups, error)
	ChatTextBot(context.Context) error
}
