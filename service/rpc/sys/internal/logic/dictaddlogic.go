package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictAddLogic) DictAdd(in *sys.DictAddReq) (*sys.DictAddResp, error) {
	// todo: add your logic here and delete this line

	return &sys.DictAddResp{}, nil
}
