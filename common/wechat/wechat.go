package wechat

//
import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

func NewMini(conf WxMiniConf) *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	return wc.GetMiniProgram(wxMiniConfig(conf))
}

//内部
func wxMiniConfig(conf WxMiniConf) *miniConfig.Config {
	memory := cache.NewMemory()
	return &miniConfig.Config{
		AppID:     conf.AppId,
		AppSecret: conf.Secret,
		Cache:     memory,
	}
}
