package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/skywalkerwei/pluton-faster/common/xerr"

	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/identity"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清除token，只针对用户服务开放访问
func (l *ClearTokenLogic) ClearToken(in *identity.ClearTokenReq) (*identity.ClearTokenResp, error) {

	userTokenKey := fmt.Sprintf("user_token:%d", in.UserId)
	if _, err := l.svcCtx.RedisClient.Del(userTokenKey); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("退出token失败"), "userId:%d,err:%v", in.UserId, err)
	}

	return &identity.ClearTokenResp{
		Ok: true,
	}, nil

	//return &identity.ClearTokenResp{}, nil
}
