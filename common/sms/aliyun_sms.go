package sms

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliSms struct {
	RegionId     string
	AccessKeyId  string
	AccessSecret string
	signName     string
}

func (s *AliSms) SendCode(template string, phone string, code string) error {
	//调用阿里云短信接口发送短信
	client, err := dysmsapi.NewClientWithAccessKey(s.RegionId, s.AccessKeyId, s.AccessSecret)

	if err != nil {
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	//request属性设置
	request.Scheme = "https"
	request.SignName = s.signName
	request.TemplateCode = template
	request.PhoneNumbers = phone
	//使用json字符串发送验证码
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	if err != nil {
		return err
	}
	//设置验证码
	request.TemplateParam = string(par)
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	//检查返回结果，并判断发生状态
	//{"RequestId":"D07C8355-1EC9-57FB-9A11-19861E18ECFB","Message":"OK","BizId":"215801540511695335^0","Code":"OK"}
	if response.Code != "OK" {
		return errors.New("短信发送不成功")
	}
	return nil
}
