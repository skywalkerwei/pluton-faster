package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

// Code2Session  code 获取 openid，SessionKey
func Code2Session(code string) (auth.ResCode2Session, error) {
	mini := NewMini()
	code2Session, err := mini.GetAuth().Code2Session(code)
	return code2Session, err
}

// Decrypt  消息解密
func Decrypt(sessionKey string, encryptedData string, iv string) (plainData *encryptor.PlainData, err error) {
	mini := NewMini()
	return mini.GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
}

// Send 发送订阅消息
func Send(msg *subscribe.Message) error {
	mini := NewMini()
	return mini.GetSubscribe().Send(msg)
}

//wechat.Send(&subscribe.Message{
//	ToUser:           "",
//	TemplateID:       "",
//	Page:             "",
//	Data:             nil,
//	MiniprogramState: "",
//	Lang:             "",
//})
