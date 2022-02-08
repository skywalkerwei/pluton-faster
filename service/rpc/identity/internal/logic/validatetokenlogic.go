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

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// validateToken ，只很对用户服务、授权服务api开放
func (l *ValidateTokenLogic) ValidateToken(in *identity.ValidateTokenReq) (*identity.ValidateTokenResp, error) {
	//userTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, in.UserId)
	userTokenKey := fmt.Sprintf("user_token:%d", in.UserId)
	dbToken, err := l.svcCtx.RedisClient.Get(userTokenKey)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR), "ValidateToken RedisClient Get userId:%d ,err:%v", in.UserId, err)
	}

	if dbToken != in.Token {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR), "ValidateToken is invalid  userId:%d , token:%s , dbToken:%s", in.UserId, in.Token, dbToken)
	}

	return &identity.ValidateTokenResp{
		Ok: true,
	}, nil

	//return &identity.ValidateTokenResp{}, nil
}
