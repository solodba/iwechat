package conf

import (
	"sync"

	"github.com/eatmoreapple/openwechat"
	"github.com/solodba/mcube/logger"
)

// 全局配置Config参数
var (
	c *Config
)

func C() *Config {
	if c == nil {
		logger.L().Panic().Msgf("please initial global config")
	}
	return c
}

// 全局配置Config结构体
type Config struct {
	App    *App    `toml:"app"`
	WeChat *WeChat `toml:"wechat"`
}

// App结构体
type App struct {
	Name string `toml:"name" env:"APP_NAME"`
}

// WeChat结构体
type WeChat struct {
	RemarkNameList []string `toml:"remark_name_list" env:"WECHAT_REMARK_NAME_LIST" envSeparator:","`
	lock           sync.Mutex
	bot            *openwechat.Bot
}

// App初始化函数
func NewDefaultApp() *App {
	return &App{
		Name: "iwechat",
	}
}

// WeChat初始化函数
func NewDefaultWeChat() *WeChat {
	return &WeChat{
		RemarkNameList: make([]string, 0),
	}
}

// Config初始化函数
func NewDefaultConfig() *Config {
	return &Config{
		App:    NewDefaultApp(),
		WeChat: NewDefaultWeChat(),
	}
}

// 初始化wechat bot
func (w *WeChat) GetWechatBot() *openwechat.Bot {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.bot = openwechat.DefaultBot(openwechat.Desktop)
	return w.bot
}
