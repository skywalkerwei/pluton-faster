package wechat

//import (
//	"ginfast/app/utils/redis_factory"
//	"github.com/silenceper/wechat/v2/miniprogram/auth"
//	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
//	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
//)
//
//// Code2Session  code 获取 openid，SessionKey
//func Code2Session(code string) (auth.ResCode2Session, error) {
//	mini := NewMini()
//	code2Session, err := mini.GetAuth().Code2Session(code)
//	key := code2Session.OpenID + "SessionKey"
//	_ = redis_factory.GetOneRedisClient().Set(key, code2Session.SessionKey, 48*60*60)
//
//	return code2Session, err
//}

// Decrypt  消息解密
//func Decrypt(openID string, encryptedData string, iv string) (plainData *encryptor.PlainData, err error) {
//	mini := NewMini()
//	key := openID + "SessionKey"
//	sessionKey, err := redis_factory.GetOneRedisClient().GetString(key)
//	if err != nil {
//		return plainData, err
//	}
//	return mini.GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
//}
//
//// Send 发送订阅消息
//func Send(msg *subscribe.Message) error {
//	mini := NewMini()
//	return mini.GetSubscribe().Send(msg)
//}

//wechat.Send(&subscribe.Message{
//	ToUser:           "",
//	TemplateID:       "",
//	Page:             "",
//	Data:             nil,
//	MiniprogramState: "",
//	Lang:             "",
//})
