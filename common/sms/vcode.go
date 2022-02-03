package sms

import (
	"bytes"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type VCode struct {
	Template    string
	Mobile      string
	Config      VCodeConfig
	RedisClient *redis.Redis
	AliSms      *AliSms
}

func NewVCode(Template string, Mobile string, Config VCodeConfig, RedisClient *redis.Redis, AliSms *AliSms) *VCode {
	return &VCode{Template, Mobile, Config, RedisClient, AliSms}
}

type VCodeConfig struct {
	Debug         bool
	Length        int
	Life          int
	MaxCheckTimes int
	MagicCode     string
	TestUsers     []string
}

// Send 发送
func (v *VCode) Send() error {
	//debug 状态不发送不校验
	if v.Config.Debug {
		return nil
	}
	//testUsers 不发送不校验
	for _, item := range v.Config.TestUsers {
		if v.Mobile == item {
			return nil
		}
	}

	key := v.getKey()
	isExists, _ := v.RedisClient.Exists(key)
	if isExists {
		return errors.New("验证码已发送，请勿重复请求")
	}
	//缓存
	code := RandCode(v.Config.Length)
	err := v.RedisClient.Setex(key, code, v.Config.Life)
	if err != nil {
		return err
	}
	//发送短信
	aliSms := AliSms{}
	return aliSms.SendCode(v.Template, v.Mobile, code)
}

// Check 验证
func (v *VCode) Check(code string) error {
	//debug 状态不发送不校验
	if v.Config.Debug {
		return nil
	}
	//testUsers 不发送不校验
	for _, item := range v.Config.TestUsers {
		if v.Mobile == item {
			return nil
		}
	}
	// 魔法密码直接放过
	if v.Config.MagicCode == code {
		return nil
	}
	// 正常校验
	key := v.getKey()

	vCode, err := v.RedisClient.Get(key)
	if err != nil {
		return errors.New("验证码有误")
	}
	if vCode != code {
		return errors.New("验证码有误")
	}
	_, _ = v.RedisClient.Del(key)
	return nil
}

func (v *VCode) getKey() string {
	return v.Template + "_validate_code_" + v.Mobile
}

//RandCode 生成随机数
func RandCode(length int) string {
	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(math.Pow10(length)))
	s := bytes.Buffer{}
	s.WriteString(strconv.Itoa(int(randNum)))
	return s.String()
}
