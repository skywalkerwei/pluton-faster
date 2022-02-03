package auth

import (
	"context"
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/wechat"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) CodeLogic {
	return CodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeLogic) Code(req types.AuthCodeRequest) (resp *types.AuthCodeCodeResponse, err error) {
	code2Session, err := wechat.Code2Session(req.Code)
	if err != nil {
		return nil, err
	}
	sessionKey := "SessionKey-" + code2Session.OpenID
	err = l.svcCtx.RedisClient.Setex(sessionKey, code2Session.SessionKey, 48*60*60)
	if err != nil {
		return nil, err
	}
	//获取用户信息
	fmt.Println("SmsConf", l.svcCtx.Config.SmsConf.MaxCheckTimes, l.svcCtx.Config.SmsConf.MagicCode)
	fmt.Println("RedisConf", l.svcCtx.Config.RedisConf.Host)
	fmt.Println("WxMiniConf", l.svcCtx.Config.WxMiniConf.AppId)

	var UserInfo types.UserInfoResponse

	return &types.AuthCodeCodeResponse{
		OpenID:   code2Session.OpenID,
		IsReg:    false,
		UserInfo: UserInfo,
	}, nil

}
