package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	// todo: add your logic here and delete this line

	return &sys.RoleListResp{}, nil
}
