package other

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendSmsLogic {
	return SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req types.SendSmsReq) (resp *types.SendSmsResp, err error) {

	//var smsConf = l.svcCtx.Config.SmsConf
	//var aliSms = sms.NewAliSms(
	//	smsConf.AliYun.RegionId,
	//	smsConf.AliYun.AccessKeyID,
	//	smsConf.AliYun.AccessKeySecret,
	//	smsConf.AliYun.SignName)
	//err = aliSms.SendCode(smsConf.Template.Reg, "18627111095", "123213")
	//if err != nil {
	//	return nil, err
	//}

	return
}
