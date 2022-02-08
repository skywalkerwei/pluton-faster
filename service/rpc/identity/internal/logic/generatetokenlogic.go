package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/skywalkerwei/pluton-faster/common/jwtx"
	"github.com/skywalkerwei/pluton-faster/common/xerr"
	"time"

	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/identity"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成token，只针对用户服务开放访问
func (l *GenerateTokenLogic) GenerateToken(in *identity.GenerateTokenReq) (*identity.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("生成token失败"), "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}

	//存入redis
	//userTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, in.UserId)
	userTokenKey := fmt.Sprintf("user_token:%d", in.UserId)
	err = l.svcCtx.RedisClient.Setex(userTokenKey, accessToken, int(accessExpire))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("生成token失败"), "SetnxEx err userId:%d, err:%v", in.UserId, err)
	}

	return &identity.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil

}
