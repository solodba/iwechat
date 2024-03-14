package protocol

// Wechat服务结构体
type WechatService struct {
}

// Wechat服务结构体初始化
func NewWechatService() *WechatService {
	return &WechatService{}
}

// Wechat服务启动方法
func (w *WechatService) Start() error {
	return nil
}

// Wechat服务停止方法
// Http服务停止方法
func (s *WechatService) Stop() error {
	return nil
}
