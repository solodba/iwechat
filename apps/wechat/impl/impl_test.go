package impl_test

import (
	"context"

	"github.com/solodba/iwechat/apps/wechat"
	"github.com/solodba/iwechat/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc wechat.Service
	ctx = context.Background()
)

// 初始化函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(wechat.AppName).(wechat.Service)
}
