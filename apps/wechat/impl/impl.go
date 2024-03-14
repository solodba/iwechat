package impl

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/solodba/iwechat/apps/wechat"
	"github.com/solodba/iwechat/conf"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	c   *conf.Config
	bot *openwechat.Bot
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return wechat.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.c = conf.C()
	i.bot = conf.C().WeChat.GetWechatBot()
	return nil
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
}
