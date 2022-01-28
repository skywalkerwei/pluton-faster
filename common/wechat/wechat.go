package wechat

//
//import (
//	"github.com/silenceper/wechat/v2"
//	"github.com/silenceper/wechat/v2/cache"
//	"github.com/silenceper/wechat/v2/miniprogram"
//	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
//)
//
//func NewMini() *miniprogram.MiniProgram {
//	wc := wechat.NewWechat()
//	return wc.GetMiniProgram(wxMiniConfig())
//}
//
////内部
//func wxMiniConfig() *miniConfig.Config {
//	//memory := wxCache()
//	memory := cache.NewMemory()
//	return &miniConfig.Config{
//		AppID:     "wxbb21d44638bd5640",
//		AppSecret: "c9e2acc9022b7d59f0743518e436e477",
//		Cache:     memory,
//	}
//}

//内部
//func wxCache() cache.Cache {
//	redisOpts := &cache.RedisOpts{
//		Host:        yml_config.CreateYamlFactory().GetString("Redis.Host") + ":" + yml_config.CreateYamlFactory().GetString("Redis.Port"),
//		Password:    yml_config.CreateYamlFactory().GetString("Redis.Auth"),
//		Database:    0,
//		MaxActive:   yml_config.CreateYamlFactory().GetInt("Redis.MaxActive"),
//		MaxIdle:     yml_config.CreateYamlFactory().GetInt("Redis.MaxIdle"),
//		IdleTimeout: yml_config.CreateYamlFactory().GetInt("Redis.IdleTimeout"),
//	}
//	return cache.NewRedis(redisOpts)
//}
