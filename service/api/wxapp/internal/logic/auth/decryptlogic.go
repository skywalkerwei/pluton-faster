package auth

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecryptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDecryptLogic(ctx context.Context, svcCtx *svc.ServiceContext) DecryptLogic {
	return DecryptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DecryptLogic) Decrypt(req types.DecryptRequest) (resp *types.DecryptResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
