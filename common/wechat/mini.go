package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

// Code2Session  code 获取 openid，SessionKey
func Code2Session(miniProgram *miniprogram.MiniProgram, code string) (auth.ResCode2Session, error) {
	code2Session, err := miniProgram.GetAuth().Code2Session(code)
	return code2Session, err
}

// Decrypt  消息解密
func Decrypt(miniProgram *miniprogram.MiniProgram, sessionKey string, encryptedData string, iv string) (plainData *encryptor.PlainData, err error) {
	return miniProgram.GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
}

// Send 发送订阅消息
func Send(miniProgram *miniprogram.MiniProgram, msg *subscribe.Message) error {
	return miniProgram.GetSubscribe().Send(msg)
}

//wechat.Send(&subscribe.Message{
//	ToUser:           "",
//	TemplateID:       "",
//	Page:             "",
//	Data:             nil,
//	MiniprogramState: "",
//	Lang:             "",
//})
