package start

import (
	"github.com/solodba/iwechat/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "iwechat start service",
	Long:    "iwechat start service",
	Example: "iwechat-api start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewServer()
		if err := srv.Start(); err != nil {
			return err
		}
		return nil
	},
}

// 服务结构体
type Server struct {
	WechatService *protocol.WechatService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		WechatService: protocol.NewWechatService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	return s.WechatService.Start()
}

// Server服务停止方法
func (s *Server) Stop() error {
	return s.WechatService.Stop()
}
