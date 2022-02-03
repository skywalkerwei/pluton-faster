package wechat

//
import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

func NewMini() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	return wc.GetMiniProgram(wxMiniConfig())
}

//内部
func wxMiniConfig() *miniConfig.Config {
	memory := cache.NewMemory()
	return &miniConfig.Config{
		AppID:     "wx5ed344b55cb695a8",
		AppSecret: "9a38893b9ec3c4f5578666b5b8781548",
		Cache:     memory,
	}
}
